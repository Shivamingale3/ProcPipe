package orchestrator

import (
	"context"
	"time"

	"github.com/Shivamingale3/ProcPipe/logger"
	"github.com/Shivamingale3/ProcPipe/monitor"
	"github.com/Shivamingale3/ProcPipe/notify"
	"github.com/Shivamingale3/ProcPipe/process"
)

func handleInput(ctx context.Context, mon *monitor.Monitor, proc *process.Process, n notify.Notifier, cmd, prompt string, isDryRun bool) {
	logger.Warn("Input detected: %s", prompt)

	// In non-dry-run mode, send prompt to Telegram
	if !isDryRun {
		logger.Info("📤 Sending prompt to Telegram...")
		n.SendInputRequired(cmd, prompt)
	}

	replyCh := make(chan inputResult, 1)
	inputCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Log the waiting state BEFORE starting input goroutines to avoid
	// a race where input arrives before the log is printed.
	if isDryRun {
		logger.Info("📩 Waiting for terminal input...")
	} else if isInteractiveTerminal() {
		logger.Info("📩 Waiting for input (terminal or Telegram)...")
	} else {
		logger.Info("📩 Waiting for Telegram reply...")
	}

	// Race: Telegram reply (only in non-dry-run mode)
	if !isDryRun {
		go func() {
			reply, err := n.WaitForReply(inputCtx)
			if err != nil {
				return
			}
			select {
			case replyCh <- inputResult{Reply: reply, Source: "telegram"}:
			default:
			}
		}()
	}

	// Race: Local stdin (always active — works for both terminal and piped input)
	go readStdin(inputCtx, replyCh)

	result := <-replyCh
	cancel() // stop the losing source

	logger.Info("📬 Received from %s: \"%s\"", result.Source, result.Reply)
	mon.Suppress()
	proc.SendInput(result.Reply + "\n")
	if !isDryRun {
		n.SendInputForwarded(result.Reply)
	}
	logger.Success("Input forwarded")
}

func handleCompletion(proc *process.Process, n notify.Notifier, mon *monitor.Monitor, cmd, host string, start time.Time) int {
	code := proc.Wait()
	dur := time.Since(start)
	if code == 0 {
		logger.Success("Completed (exit 0) in %s", dur.Round(time.Second))
	} else {
		logger.Error("Failed (exit %d) in %s", code, dur.Round(time.Second))
	}
	logger.Info("📤 Sending completion report...")
	n.SendCompleted(notify.CompleteInfo{
		Command: cmd, ExitCode: code, Duration: dur,
		Logs: mon.Logs(), Host: host,
	})
	logger.Success("📡 Telegram notified")
	return code
}
