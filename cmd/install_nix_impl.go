package cmd

import (
	"os"
	"os/exec"

	"github.com/Shivamingale3/ProcPipe/logger"
)

func installUnix(self string) {
	dest := "/usr/local/bin/procpipe"
	if self == dest {
		logger.Success("Already installed at %s", dest)
		return
	}
	logger.Info("Installing to %s ...", dest)
	if err := tryCopy(self, dest); err != nil {
		logger.Info("Need sudo for /usr/local/bin")
		cmd := exec.Command("sudo", "cp", self, dest)
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
		if err := cmd.Run(); err != nil {
			logger.Error("Install failed: %s", err)
			return
		}
		exec.Command("sudo", "chmod", "+x", dest).Run()
	}
	logger.Success("Installed! Run 'procpipe' from anywhere.")
	done()
}
