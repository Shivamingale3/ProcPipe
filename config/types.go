package config

// TelegramConfig holds Telegram bot credentials.
type TelegramConfig struct {
	BotToken string `yaml:"bot_token"`
	ChatID   int64  `yaml:"chat_id"`
}

// MonitorConfig holds output monitoring settings.
type MonitorConfig struct {
	LogTailLines  int      `yaml:"log_tail_lines"`
	InputPatterns []string `yaml:"input_patterns"`
}

// Config is the top-level application configuration.
type Config struct {
	Telegram TelegramConfig `yaml:"telegram"`
	Monitor  MonitorConfig  `yaml:"monitor"`
	Command  []string       `yaml:"-"`
	DryRun   bool           `yaml:"-"`
}
