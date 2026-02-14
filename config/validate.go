package config

import "fmt"

// validate ensures all required config fields are present.
func validate(cfg *Config) error {
	if !cfg.DryRun {
		if cfg.Telegram.BotToken == "" {
			return fmt.Errorf("telegram.bot_token required (config or --token)")
		}
		if cfg.Telegram.ChatID == 0 {
			return fmt.Errorf("telegram.chat_id required (config or --chat)")
		}
	}
	if len(cfg.Command) == 0 {
		return fmt.Errorf("no command given. Usage: procpipe [flags] -- <command>")
	}
	return nil
}
