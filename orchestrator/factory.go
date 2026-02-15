package orchestrator

import (
	"os"

	"github.com/Shivamingale3/ProcPipe/config"
	"github.com/Shivamingale3/ProcPipe/notify"
	"github.com/Shivamingale3/ProcPipe/telegram"
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
