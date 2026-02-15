package cmd

import (
	"fmt"

	"procpipe/config"

	"github.com/spf13/cobra"
)

var configPathCmd = &cobra.Command{
	Use:   "path",
	Short: "Print config file location",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println(config.ConfigPath(""))
	},
}

func init() {
	configCmd.AddCommand(configPathCmd)
}
