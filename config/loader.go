package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Load reads config from YAML, overlays CLI flags, and validates.
func Load(flags *Flags) (*Config, error) {
	cfg := DefaultConfig()

	path := resolveConfigPath(flags.ConfigFile)
	if data, err := os.ReadFile(path); err == nil {
		if err := yaml.Unmarshal(data, cfg); err != nil {
			return nil, fmt.Errorf("parsing config %s: %w", path, err)
		}
	}

	applyFlags(cfg, flags)
	return cfg, validate(cfg)
}

func resolveConfigPath(custom string) string {
	if custom != "" {
		return custom
	}
	if _, err := os.Stat("procpipe.yaml"); err == nil {
		return "procpipe.yaml"
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".procpipe.yaml")
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
