package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Shivamingale3/ProcPipe/config"
	"github.com/Shivamingale3/ProcPipe/logger"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure ProcPipe (interactive setup)",
	Run:   configWizard,
}

func init() { rootCmd.AddCommand(configCmd) }

func configWizard(_ *cobra.Command, _ []string) {
	bold := color.New(color.Bold).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(bold("🔧 ProcPipe Configuration"))
	fmt.Println(cyan("━━━━━━━━━━━━━━━━━━━━━━━━━"))
	fmt.Println(bold("📖 Step 1: Create a Telegram bot"))
	fmt.Println("   → Open Telegram, search @BotFather")
	fmt.Println("   → Send /newbot and follow the prompts")
	fmt.Println("   → Copy the bot token")

	token := prompt(reader, "Enter Bot Token")
	chatID := promptInt(reader, "Enter Chat ID")

	cfg := config.DefaultConfig()
	cfg.Telegram.BotToken = token
	cfg.Telegram.ChatID = chatID

	if err := config.Save(cfg); err != nil {
		logger.Error("Failed to save: %s", err)
		return
	}
	logger.Success("Config saved to %s", config.ConfigPath(""))
}
