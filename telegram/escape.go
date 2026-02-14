package telegram

import "strings"

// escapeHTML escapes special HTML characters for Telegram's HTML parse mode.
func escapeHTML(s string) string {
	r := strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;")
	return r.Replace(s)
}

// truncateLogs trims logs to fit within Telegram's 4096 char message limit.
func truncateLogs(logs string, maxLen int) string {
	if len(logs) <= maxLen {
		return logs
	}
	return "…(truncated)\n" + logs[len(logs)-maxLen:]
}
