package notify

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
)

type DryRun struct{}

func NewDryRun() *DryRun { return &DryRun{} }
func (d *DryRun) SendStarted(info StartInfo) error {
	color.Green("🚀 Process Started: %s", info.Command)
	color.White("   Host: %s | Dir: %s", info.Host, info.Directory)
	return nil
}
func (d *DryRun) SendCompleted(info CompleteInfo) error {
	if info.ExitCode == 0 {
		color.Green("✅ Completed (exit 0) in %s", info.Duration)
	} else {
		color.Red("❌ Failed (exit %d) in %s", info.ExitCode, info.Duration)
	}
	fmt.Println("━━━ Output ━━━")
	fmt.Println(info.Logs)
	return nil
}
func (d *DryRun) SendInputRequired(cmd, prompt string) error {
	color.Yellow("⚠️  Input required: %s", prompt)
	fmt.Print("Enter input: ")
	return nil
}
func (d *DryRun) SendInputForwarded(input string) error {
	color.Cyan("📨 Input forwarded: %s", input)
	return nil
}
func (d *DryRun) WaitForReply(_ context.Context) (string, error) {
	s := bufio.NewScanner(os.Stdin)
	if s.Scan() {
		return s.Text(), nil
	}
	return "", fmt.Errorf("no input")
}
