package orchestrator

import (
	"context"
	"time"

	"procpipe/logger"
	"procpipe/monitor"
	"procpipe/notify"
	"procpipe/process"
)

func handleEvents(mon *monitor.Monitor, proc *process.Process, n notify.Notifier, cmd, host string, start time.Time) int {
	ctx := context.Background()
	for event := range mon.Events() {
		switch event.Type {
		case monitor.EventInputPrompt:
			handleInput(ctx, proc, n, cmd, event.Line)
		case monitor.EventProcessDone:
			return handleCompletion(proc, n, mon, cmd, host, start)
		}
	}
	return proc.Wait()
}

func handleInput(ctx context.Context, proc *process.Process, n notify.Notifier, cmd, prompt string) {
	logger.Warn("Input detected: %s", prompt)
	n.SendInputRequired(cmd, prompt)
	reply, err := n.WaitForReply(ctx)
	if err != nil {
		logger.Error("Failed to get reply: %s", err)
		return
	}
	proc.SendInput(reply + "\n")
	n.SendInputForwarded(reply)
	logger.Success("Input forwarded: %s", reply)
}

func handleCompletion(proc *process.Process, n notify.Notifier, mon *monitor.Monitor, cmd, host string, start time.Time) int {
	code := proc.Wait()
	dur := time.Since(start)
	if code == 0 {
		logger.Success("Process completed (exit 0) in %s", dur.Round(time.Second))
	} else {
		logger.Error("Process failed (exit %d) in %s", code, dur.Round(time.Second))
	}
	n.SendCompleted(notify.CompleteInfo{
		Command: cmd, ExitCode: code, Duration: dur,
		Logs: mon.Logs(), Host: host,
	})
	return code
}

