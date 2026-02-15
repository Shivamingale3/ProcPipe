package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Load(flags *Flags) (*Config, error) {
	cfg := DefaultConfig()

	path := ConfigPath(flags.ConfigFile)
	if data, err := os.ReadFile(path); err == nil {
		if err := yaml.Unmarshal(data, cfg); err != nil {
			return nil, fmt.Errorf("parsing config %s: %w", path, err)
		}
	}

	applyFlags(cfg, flags)
	return cfg, validate(cfg)
}

func LoadPartial() (*Config, error) {
	cfg := DefaultConfig()
	path := ConfigPath("")
	if data, err := os.ReadFile(path); err == nil {
		if err := yaml.Unmarshal(data, cfg); err != nil {
			return nil, fmt.Errorf("parsing config %s: %w", path, err)
		}
	}
	return cfg, validatePartial(cfg)
}

func applyFlags(cfg *Config, f *Flags) {
	if f.BotToken != "" {
		cfg.Telegram.BotToken = f.BotToken
	}
	if f.ChatID != 0 {
		cfg.Telegram.ChatID = f.ChatID
	}
	cfg.DryRun = f.DryRun
	cfg.Command = f.Command
}

