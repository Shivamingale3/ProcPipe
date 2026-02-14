package config

import "flag"

// Flags holds CLI flag values parsed before config loading.
type Flags struct {
	BotToken   string
	ChatID     int64
	DryRun     bool
	ConfigFile string
	Command    []string
}

// ParseFlags parses CLI arguments and returns a Flags struct.
func ParseFlags() *Flags {
	f := &Flags{}
	flag.StringVar(&f.BotToken, "token", "", "Telegram bot token")
	flag.Int64Var(&f.ChatID, "chat", 0, "Telegram chat ID")
	flag.BoolVar(&f.DryRun, "dry-run", false, "Print notifications to stdout")
	flag.StringVar(&f.ConfigFile, "config", "", "Path to config file")
	flag.Parse()
	f.Command = flag.Args()
	return f
}
