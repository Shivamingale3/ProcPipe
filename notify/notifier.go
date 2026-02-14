package notify

import (
	"context"
	"time"
)

// StartInfo carries metadata about a freshly started process.
type StartInfo struct {
	Command   string
	Host      string
	Directory string
	StartTime time.Time
}

// CompleteInfo carries metadata about a finished process.
type CompleteInfo struct {
	Command  string
	ExitCode int
	Duration time.Duration
	Logs     string
	Host     string
}

// Notifier is the interface for sending process status notifications.
type Notifier interface {
	SendStarted(info StartInfo) error
	SendCompleted(info CompleteInfo) error
	SendInputRequired(command, promptLine string) error
	SendInputForwarded(input string) error
	WaitForReply(ctx context.Context) (string, error)
}
