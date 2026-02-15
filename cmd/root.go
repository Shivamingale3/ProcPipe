package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "procpipe",
	Short: "Terminal process watcher with Telegram notifications",
	Long:  "ProcPipe watches long-running commands and notifies you on Telegram when they complete or need input.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
