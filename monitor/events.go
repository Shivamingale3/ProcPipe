package monitor

// EventType identifies what kind of monitor event occurred.
type EventType int

const (
	EventOutput      EventType = iota // A complete line of output
	EventInputPrompt                  // An input prompt was detected
	EventProcessDone                  // The process finished (PTY EOF)
)

// Event carries data from the output monitor to the orchestrator.
type Event struct {
	Type   EventType
	Line   string // The output line
	Prompt string // Matched pattern (for InputPrompt events)
}
