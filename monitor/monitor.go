package monitor

import (
	"io"
	"sync"
)

type Monitor struct {
	reader     io.Reader
	buffer     *RingBuffer
	matcher    *Matcher
	events     chan Event
	lastPrompt string
	suppressed bool
	mu         sync.Mutex
}

func New(reader io.Reader, bufSize int, customPatterns []string) *Monitor {
	patterns := DefaultPatterns
	if len(customPatterns) > 0 {
		patterns = append(patterns, customPatterns...)
	}
	return &Monitor{
		reader: reader, buffer: NewRingBuffer(bufSize),
		matcher: NewMatcher(CompilePatterns(patterns)),
		events:  make(chan Event, 16),
	}
}

func (m *Monitor) Events() <-chan Event { return m.events }
func (m *Monitor) Logs() string         { return m.buffer.String() }
