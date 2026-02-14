package monitor

import (
	"fmt"
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

func (m *Monitor) Start() {
	go func() {
		ReadOutput(m.reader, m.handleOutput)
		m.events <- Event{Type: EventProcessDone}
		close(m.events)
	}()
}

func (m *Monitor) Events() <-chan Event { return m.events }
func (m *Monitor) Logs() string         { return m.buffer.String() }

// Suppress stops prompt detection until the next fresh output line.
func (m *Monitor) Suppress() {
	m.mu.Lock()
	m.suppressed = true
	m.mu.Unlock()
}

func (m *Monitor) handleOutput(line string, isPartial bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if !isPartial {
		m.buffer.Add(line)
		fmt.Println(line)
		m.suppressed = false
		m.lastPrompt = ""
		return
	}
	if m.suppressed {
		return
	}
	if match, prompt := m.matcher.IsInputPrompt(line); match {
		if prompt != m.lastPrompt {
			m.lastPrompt = prompt
			m.events <- Event{Type: EventInputPrompt, Prompt: prompt, Line: line}
		}
	}
}

