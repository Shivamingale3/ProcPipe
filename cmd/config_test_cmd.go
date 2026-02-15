package cmd

import (
	"procpipe/config"
	"procpipe/logger"
	"procpipe/telegram"

	"github.com/spf13/cobra"
)

var configTestCmd = &cobra.Command{
	Use:   "test",
	Short: "Test Telegram bot connection",
	Run: func(_ *cobra.Command, _ []string) {
		cfg, err := config.LoadPartial()
		if err != nil {
			logger.Error("No config: %s", err)
			return
		}
		logger.Info("Testing Telegram connection...")
		client := telegram.NewClient(cfg.Telegram.BotToken, cfg.Telegram.ChatID)
		if err := client.SendMessage("🔔 <b>ProcPipe Test</b>\n\nConnection successful!"); err != nil {
			logger.Error("Failed: %s", err)
			return
		}
		logger.Success("Message sent! Check your Telegram.")
	},
}

func init() {
	configCmd.AddCommand(configTestCmd)
}
