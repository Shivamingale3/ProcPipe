package orchestrator

import (
	"bufio"
	"context"
	"os"

	"golang.org/x/term"
)

// inputResult carries a reply from either stdin or Telegram.
type inputResult struct {
	Reply  string
	Source string // "terminal" or "telegram"
}

// isInteractiveTerminal returns true if stdin is a real terminal (not piped).
func isInteractiveTerminal() bool {
	return term.IsTerminal(int(os.Stdin.Fd()))
}

// readStdin reads one line from stdin in a goroutine.
// Sends the result to ch. Respects context cancellation.
func readStdin(ctx context.Context, ch chan<- inputResult) {
	lineCh := make(chan string, 1)

	go func() {
		s := bufio.NewScanner(os.Stdin)
		if s.Scan() {
			lineCh <- s.Text()
		}
	}()

	select {
	case <-ctx.Done():
		return
	case line := <-lineCh:
		select {
		case ch <- inputResult{Reply: line, Source: "terminal"}:
		default: // channel full — another source already won
		}
	}
}
