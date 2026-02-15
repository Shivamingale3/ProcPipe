package cmd

import (
	"fmt"

	"github.com/Shivamingale3/ProcPipe/config"
	"github.com/Shivamingale3/ProcPipe/logger"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current configuration",
	Run: func(_ *cobra.Command, _ []string) {
		cfg, err := config.LoadPartial()
		if err != nil {
			logger.Error("No config found: %s", err)
			return
		}
		bold := color.New(color.Bold).SprintFunc()
		token := cfg.Telegram.BotToken
		if len(token) > 10 {
			token = token[:10] + "..." // redact
		}
		fmt.Printf("%s %s\n", bold("Bot Token:"), token)
		fmt.Printf("%s %d\n", bold("Chat ID:"), cfg.Telegram.ChatID)
		fmt.Printf("%s %d\n", bold("Log Lines:"), cfg.Monitor.LogTailLines)
	},
}

func init() {
	configCmd.AddCommand(configShowCmd)
}
