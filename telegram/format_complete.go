package telegram

import (
	"fmt"
	"time"

	"procpipe/notify"
)

func formatCompleted(info notify.CompleteInfo) string {
	emoji := "✅"
	status := "Completed Successfully"
	if info.ExitCode != 0 {
		emoji = "❌"
		status = "Failed"
	}

	logs := truncateLogs(info.Logs, 3000)
	dur := info.Duration.Round(time.Second).String()

	return fmt.Sprintf(
		"%s <b>Process %s</b>\n\n"+
			"📋 <b>Command:</b>\n<code>%s</code>\n\n"+
			"📊 <b>Exit Code:</b> <code>%d</code>\n"+
			"⏱️ <b>Duration:</b> %s\n"+
			"🖥️ <b>Host:</b> %s\n\n"+
			"━━━ 📝 Output (last lines) ━━━\n"+
			"<pre>%s</pre>",
		emoji, status,
		escapeHTML(info.Command),
		info.ExitCode,
		dur,
		escapeHTML(info.Host),
		escapeHTML(logs),
	)
}

func formatInputForwarded(input string) string {
	return fmt.Sprintf(
		"📨 <b>Input Forwarded</b>\n\n"+
			"Your response <code>%s</code> has been sent to the process.\n\n"+
			"👁️ Continuing to watch...",
		escapeHTML(input),
	)
}
