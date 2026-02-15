package orchestrator

import (
	"context"
	"time"

	"github.com/Shivamingale3/ProcPipe/logger"
	"github.com/Shivamingale3/ProcPipe/monitor"
	"github.com/Shivamingale3/ProcPipe/notify"
	"github.com/Shivamingale3/ProcPipe/process"
)

func handleInput(ctx context.Context, mon *monitor.Monitor, proc *process.Process, n notify.Notifier, cmd, prompt string) {
	logger.Warn("Input detected: %s", prompt)
	logger.Info("📤 Sending prompt to Telegram...")
	n.SendInputRequired(cmd, prompt)
	logger.Info("📩 Waiting for Telegram reply...")
	reply, err := n.WaitForReply(ctx)
	if err != nil {
		logger.Error("Failed to get reply: %s", err)
		return
	}
	logger.Info("📬 Received: \"%s\"", reply)
	mon.Suppress()
	proc.SendInput(reply + "\n")
	n.SendInputForwarded(reply)
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
