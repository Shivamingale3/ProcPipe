package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/Shivamingale3/ProcPipe/logger"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install ProcPipe to system PATH",
	Run:   installSelf,
}

func init() { rootCmd.AddCommand(installCmd) }

func installSelf(_ *cobra.Command, _ []string) {
	self, err := os.Executable()
	if err != nil {
		logger.Error("Cannot find executable: %s", err)
		return
	}
	self, _ = filepath.EvalSymlinks(self)
	switch runtime.GOOS {
	case "linux", "darwin":
		installUnix(self)
	case "windows":
		installWindows(self)
	default:
		logger.Error("Unsupported OS: %s", runtime.GOOS)
	}
}

func tryCopy(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0755)
}

func done() { fmt.Println("  Try: procpipe version") }
