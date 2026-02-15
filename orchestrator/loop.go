package orchestrator

import (
	"context"
	"time"

	"github.com/Shivamingale3/ProcPipe/monitor"
	"github.com/Shivamingale3/ProcPipe/notify"
	"github.com/Shivamingale3/ProcPipe/process"
)

func handleEvents(mon *monitor.Monitor, proc *process.Process, n notify.Notifier, cmd, host string, start time.Time) int {
	ctx := context.Background()
	for event := range mon.Events() {
		switch event.Type {
		case monitor.EventInputPrompt:
			handleInput(ctx, mon, proc, n, cmd, event.Line)
		case monitor.EventProcessDone:
			return handleCompletion(proc, n, mon, cmd, host, start)
		}
	}
	return proc.Wait()
}
