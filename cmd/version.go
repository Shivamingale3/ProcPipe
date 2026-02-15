package cmd

import (
	"fmt"
	"runtime"

	"procpipe/version"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show ProcPipe version and build info",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("ProcPipe v%s\n", version.Version)
		fmt.Printf("Commit:  %s\n", version.Commit)
		fmt.Printf("Built:   %s\n", version.BuildDate)
		fmt.Printf("OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
		fmt.Printf("Go:      %s\n", runtime.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
