package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Shivamingale3/ProcPipe/logger"
)

func installWindows(self string) {
	dir := filepath.Join(os.Getenv("LOCALAPPDATA"), "ProcPipe")
	dest := filepath.Join(dir, "procpipe.exe")
	os.MkdirAll(dir, 0755)
	logger.Info("Installing to %s ...", dir)
	if err := tryCopy(self, dest); err != nil {
		logger.Error("Copy failed: %s", err)
		return
	}
	ps := fmt.Sprintf(
		`$p = [Environment]::GetEnvironmentVariable('PATH','User'); `+
			`if ($p -notlike '*%s*') { `+
			`[Environment]::SetEnvironmentVariable('PATH', $p + ';%s', 'User') }`,
		dir, dir,
	)
	cmd := exec.Command("powershell", "-Command", ps)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		logger.Warn("Could not add to PATH: %s", err)
		logger.Info("Add manually: %s", dir)
		return
	}
	logger.Success("Installed! Restart terminal, then run 'procpipe'")
	done()
}
