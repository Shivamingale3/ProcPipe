package process

import (
	"os"
	"os/exec"
)

// Process wraps a PTY-spawned command with I/O handles.
type Process struct {
	cmd      *exec.Cmd
	pty      *os.File
	done     chan struct{}
	exitCode int
}

// Done returns a channel that closes when the process exits.
func (p *Process) Done() <-chan struct{} { return p.done }

// ExitCode returns the exit code after the process exits.
func (p *Process) ExitCode() int { return p.exitCode }
