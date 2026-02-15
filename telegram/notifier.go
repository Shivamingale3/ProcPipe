package telegram

import (
	"context"
	"github.com/Shivamingale3/ProcPipe/notify"
)

// Notifier implements notify.Notifier using the Telegram Bot API.
type Notifier struct {
	client *Client
}

// NewNotifier creates a Telegram-backed notifier.
func NewNotifier(token string, chatID int64) *Notifier {
	return &Notifier{client: NewClient(token, chatID)}
}

func (t *Notifier) SendStarted(info notify.StartInfo) error {
	return t.client.SendMessage(formatStarted(info))
}

func (t *Notifier) SendCompleted(info notify.CompleteInfo) error {
	return t.client.SendMessage(formatCompleted(info))
}

func (t *Notifier) SendInputRequired(cmd, prompt string) error {
	return t.client.SendMessage(formatInputRequired(cmd, prompt))
}

func (t *Notifier) SendInputForwarded(input string) error {
	return t.client.SendMessage(formatInputForwarded(input))
}

func (t *Notifier) WaitForReply(ctx context.Context) (string, error) {
	return t.client.PollForReply(ctx)
}
