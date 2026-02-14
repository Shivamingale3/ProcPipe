package monitor

import "strings"

// RingBuffer stores the last N lines of output efficiently.
type RingBuffer struct {
	lines []string
	size  int
	pos   int
	full  bool
}

// NewRingBuffer creates a ring buffer of the given capacity.
func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{lines: make([]string, size), size: size}
}

// Add appends a line, overwriting the oldest if full.
func (rb *RingBuffer) Add(line string) {
	rb.lines[rb.pos] = line
	rb.pos = (rb.pos + 1) % rb.size
	if rb.pos == 0 {
		rb.full = true
	}
}

// String returns all stored lines in chronological order.
func (rb *RingBuffer) String() string {
	if !rb.full {
		return strings.Join(rb.lines[:rb.pos], "\n")
	}
	ordered := make([]string, 0, rb.size)
	ordered = append(ordered, rb.lines[rb.pos:]...)
	ordered = append(ordered, rb.lines[:rb.pos]...)
	return strings.Join(ordered, "\n")
}
