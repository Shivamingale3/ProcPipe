package telegram

import (
	"fmt"
	"procpipe/notify"
)

func formatStarted(info notify.StartInfo) string {
	return fmt.Sprintf(
		"🚀 <b>Process Started</b>\n\n"+
			"📋 <b>Command:</b>\n<code>%s</code>\n\n"+
			"🖥️ <b>Host:</b> %s\n"+
			"📁 <b>Directory:</b> <code>%s</code>\n"+
			"🕐 <b>Started:</b> %s\n\n"+
			"━━━━━━━━━━━━━━━━━━━━━━━\n"+
			"👁️ Watching... I'll notify you on completion or if input is needed.",
		escapeHTML(info.Command),
		escapeHTML(info.Host),
		escapeHTML(info.Directory),
		info.StartTime.Format("2006-01-02 15:04:05"),
	)
}

func formatInputRequired(command, promptLine string) string {
	return fmt.Sprintf(
		"⚠️ <b>Input Required</b>\n\n"+
			"📋 <b>Command:</b>\n<code>%s</code>\n\n"+
			"🔔 <b>Prompt Detected:</b>\n"+
			"<pre>%s</pre>\n\n"+
			"━━━━━━━━━━━━━━━━━━━━━━━\n"+
			"💬 <b>Reply to this message with your input.</b>",
		escapeHTML(command),
		escapeHTML(promptLine),
	)
}
