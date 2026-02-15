package monitor

import "fmt"

func (m *Monitor) Start() {
	go func() {
		ReadOutput(m.reader, m.handleOutput)
		m.events <- Event{Type: EventProcessDone}
		close(m.events)
	}()
}

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
