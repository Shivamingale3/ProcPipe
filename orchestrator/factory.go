package orchestrator

import (
	"os"

	"procpipe/config"
	"procpipe/notify"
	"procpipe/telegram"
)

func createNotifier(cfg *config.Config) notify.Notifier {
	if cfg.DryRun {
		return notify.NewDryRun()
	}
	return telegram.NewNotifier(cfg.Telegram.BotToken, cfg.Telegram.ChatID)
}

func mustGetwd() string {
	dir, _ := os.Getwd()
	return dir
}
