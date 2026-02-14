package monitor

import "regexp"

// Matcher checks lines against compiled input prompt patterns.
type Matcher struct {
	patterns []*regexp.Regexp
}

// NewMatcher creates a matcher from compiled regex patterns.
func NewMatcher(patterns []*regexp.Regexp) *Matcher {
	return &Matcher{patterns: patterns}
}

// IsInputPrompt returns true if the line matches any input pattern.
func (m *Matcher) IsInputPrompt(line string) (bool, string) {
	for _, re := range m.patterns {
		if re.MatchString(line) {
			return true, re.String()
		}
	}
	return false, ""
}
