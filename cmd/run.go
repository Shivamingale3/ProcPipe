package cmd

import (
	"procpipe/config"
	"procpipe/logger"
	"procpipe/orchestrator"

	"github.com/spf13/cobra"
)

var (
	flagToken  string
	flagChatID int64
	flagDryRun bool
	flagConfig string
)

var runCmd = &cobra.Command{
	Use:   "run -- <command> [args...]",
	Short: "Run and watch a command",
	Long:  "Spawns a command in a PTY, watches output, and notifies via Telegram.",
	Args:  cobra.MinimumNArgs(1),
	Run:   runCommand,
}

func init() {
	runCmd.Flags().StringVar(&flagToken, "token", "", "Telegram bot token")
	runCmd.Flags().Int64Var(&flagChatID, "chat", 0, "Telegram chat ID")
	runCmd.Flags().BoolVar(&flagDryRun, "dry-run", false, "Print to stdout")
	runCmd.Flags().StringVar(&flagConfig, "config", "", "Config file path")
	rootCmd.AddCommand(runCmd)
}

func runCommand(_ *cobra.Command, args []string) {
	cfg, err := config.Load(&config.Flags{
		BotToken: flagToken, ChatID: flagChatID,
		DryRun: flagDryRun, ConfigFile: flagConfig, Command: args,
	})
	if err != nil {
		logger.Error("Config: %s", err)
		return
	}
	if code, err := orchestrator.Run(cfg); err != nil {
		logger.Error("Fatal: %s", err)
	} else if code != 0 {
		logger.Error("Exit: %d", code)
	}
}
