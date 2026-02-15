package cmd

import (
	"fmt"
	"runtime"

	"github.com/Shivamingale3/ProcPipe/version"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show ProcPipe version and build info",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("ProcPipe %s\n", version.Version)
		fmt.Printf("Commit:  %s\n", version.Commit)
		fmt.Printf("Built:   %s\n", version.BuildDate)
		fmt.Printf("OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
		fmt.Printf("Go:      %s\n", runtime.Version())
		fmt.Printf("Developer:      %s\n", version.Developer)
		fmt.Printf("Email:      %s\n", version.Email)
		fmt.Printf("Github:      %s\n", version.Github)
		fmt.Printf("Website:      %s\n", version.Website)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
