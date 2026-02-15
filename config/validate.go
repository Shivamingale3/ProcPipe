package config

import "fmt"

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
		return fmt.Errorf("no command given. Usage: procpipe run -- <command>")
	}
	return nil
}

func validatePartial(cfg *Config) error {
	if cfg.Telegram.BotToken == "" {
		return fmt.Errorf("telegram.bot_token not set. Run: procpipe config")
	}
	if cfg.Telegram.ChatID == 0 {
		return fmt.Errorf("telegram.chat_id not set. Run: procpipe config")
	}
	return nil
}
