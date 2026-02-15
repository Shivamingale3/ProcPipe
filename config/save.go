package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func ConfigPath(custom string) string {
	if custom != "" {
		return custom
	}
	if _, err := os.Stat("procpipe.yaml"); err == nil {
		return "procpipe.yaml"
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".procpipe.yaml")
}

func Save(cfg *Config) error {
	path := ConfigPath("")
	home, _ := os.UserHomeDir()
	path = filepath.Join(home, ".procpipe.yaml")

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}
	return os.WriteFile(path, data, 0600)
}
