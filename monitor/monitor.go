package monitor

import (
	"fmt"
	"io"
)

type Monitor struct {
	reader     io.Reader
	buffer     *RingBuffer
	matcher    *Matcher
	events     chan Event
	lastPrompt string
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

func (m *Monitor) Start() {
	go func() {
		ReadOutput(m.reader, m.handleOutput)
		m.events <- Event{Type: EventProcessDone}
		close(m.events)
	}()
}

func (m *Monitor) Events() <-chan Event { return m.events }
func (m *Monitor) Logs() string         { return m.buffer.String() }

func (m *Monitor) handleOutput(line string, isPartial bool) {
	if !isPartial {
		m.buffer.Add(line)
		fmt.Println(line)
		m.lastPrompt = ""
	}
	if match, pattern := m.matcher.IsInputPrompt(line); match {
		if pattern != m.lastPrompt {
			m.lastPrompt = pattern
			m.events <- Event{Type: EventInputPrompt, Prompt: pattern, Line: line}
		}
	}
}
