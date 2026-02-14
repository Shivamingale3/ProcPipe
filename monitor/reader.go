package monitor

import (
	"io"
	"strings"
)

// LineHandler is called for each chunk of output. isPartial indicates
// the line hasn't ended yet (e.g. an input prompt with no trailing newline).
type LineHandler func(line string, isPartial bool)

// ReadOutput reads from r in raw chunks, splits into lines, and calls
// handler for both complete and partial lines. Blocks until EOF.
func ReadOutput(r io.Reader, handler LineHandler) {
	buf := make([]byte, 4096)
	var partial strings.Builder

	for {
		n, err := r.Read(buf)
		if n > 0 {
			processChunk(&partial, string(buf[:n]), handler)
		}
		if err != nil {
			if partial.Len() > 0 {
				handler(partial.String(), false)
			}
			return
		}
	}
}

func processChunk(partial *strings.Builder, chunk string, h LineHandler) {
	partial.WriteString(chunk)
	content := partial.String()
	lines := strings.Split(content, "\n")

	for _, line := range lines[:len(lines)-1] {
		h(strings.TrimRight(line, "\r"), false)
	}

	partial.Reset()
	last := lines[len(lines)-1]
	partial.WriteString(last)

	if len(last) > 0 {
		h(last, true)
	}
}
