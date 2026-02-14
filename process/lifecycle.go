package process

import (
	"io"
	"os/exec"
)

// Output returns a reader for the PTY output (blocking reads).
func (p *Process) Output() io.Reader { return p.pty }

// SendInput writes text to the process's PTY stdin.
func (p *Process) SendInput(text string) error {
	_, err := p.pty.Write([]byte(text))
	return err
}

// Wait blocks until the process exits and returns the exit code.
func (p *Process) Wait() int {
	<-p.done
	return p.exitCode
}

// Close releases the PTY file descriptor.
func (p *Process) Close() { p.pty.Close() }

// waitForExit runs in a goroutine, waits for cmd to finish.
func (p *Process) waitForExit() {
	err := p.cmd.Wait()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			p.exitCode = exitErr.ExitCode()
		} else {
			p.exitCode = -1
		}
	}
	close(p.done)
}
