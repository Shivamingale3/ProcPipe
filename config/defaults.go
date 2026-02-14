package config

const defaultLogTailLines = 50

// DefaultConfig returns a Config with sensible defaults.
func DefaultConfig() *Config {
	return &Config{
		Monitor: MonitorConfig{
			LogTailLines: defaultLogTailLines,
		},
	}
}
