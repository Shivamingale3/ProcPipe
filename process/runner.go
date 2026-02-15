package process

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/creack/pty"
)

// Start spawns the command inside a pseudo-terminal.
func Start(command []string) (*Process, error) {
	if len(command) == 0 {
		return nil, fmt.Errorf("empty command")
	}

	if len(command) == 1 && strings.ContainsAny(command[0], "&|;<>") {
		command = []string{"/bin/sh", "-c", command[0]}
	}

	cmd := exec.Command(command[0], command[1:]...)
	ptmx, err := pty.Start(cmd)
	if err != nil {
		return nil, fmt.Errorf("pty start: %w", err)
	}

	proc := &Process{
		cmd:  cmd,
		pty:  ptmx,
		done: make(chan struct{}),
	}

	go proc.waitForExit()
	return proc, nil
}
