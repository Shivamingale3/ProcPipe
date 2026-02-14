package monitor

import "regexp"

// DefaultPatterns are regex patterns that detect interactive input prompts.
var DefaultPatterns = []string{
	`(?i)\[y/n\]`,
	`(?i)\[yes/no\]`,
	`(?i)password\s*:\s*$`,
	`(?i)passphrase.*:\s*$`,
	`(?i)enter\s+.*:\s*$`,
	`(?i)\(y/n\)`,
	`(?i)continue\s*\?\s*`,
	`(?i)press\s+(enter|return|any\s+key)`,
	`(?i)type\s+.*to\s+confirm`,
	`(?i)do you want to`,
	`(?i)are you sure`,
	`(?i)\[sudo\]\s*password`,
}

// CompilePatterns compiles a list of regex strings, skipping invalid ones.
func CompilePatterns(patterns []string) []*regexp.Regexp {
	compiled := make([]*regexp.Regexp, 0, len(patterns))
	for _, p := range patterns {
		if re, err := regexp.Compile(p); err == nil {
			compiled = append(compiled, re)
		}
	}
	return compiled
}
