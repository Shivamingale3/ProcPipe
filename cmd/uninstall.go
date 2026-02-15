package cmd

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"procpipe/logger"

	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall ProcPipe from system",
	Run:   uninstallSelf,
}

func init() { rootCmd.AddCommand(uninstallCmd) }

func uninstallSelf(_ *cobra.Command, _ []string) {
	switch runtime.GOOS {
	case "linux", "darwin":
		uninstallUnix()
	case "windows":
		uninstallWindows()
	default:
		logger.Error("Unsupported OS: %s", runtime.GOOS)
	}
}

func uninstallUnix() {
	dest := "/usr/local/bin/procpipe"
	logger.Info("Removing %s ...", dest)
	cmd := exec.Command("sudo", "rm", dest)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		logger.Error("Uninstall failed: %s", err)
		return
	}
	logger.Success("Uninstalled successfully.")
}

func uninstallWindows() {
	dir := filepath.Join(os.Getenv("LOCALAPPDATA"), "ProcPipe")
	logger.Info("Removing %s ...", dir)
	// Remove from PATH would be nice but complex via code
	cmd := exec.Command("powershell", "-Command", "Remove-Item -Recurse -Force '"+dir+"'")
	if err := cmd.Run(); err != nil {
		logger.Error("Failed to remove dir: %s", err)
	}
	logger.Success("Uninstalled. Note: Remove from PATH manually.")
}
