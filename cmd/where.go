package cmd

import (
	"fmt"

	"github.com/kevinliao852/smart-commit/config"
	"github.com/spf13/cobra"
)

var whereCmd = &cobra.Command{
	Use:   "where",
	Short: "Show the configuration file path",
	Long:  `Show the path of the configuration file being used by smart-commit.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Get()
		if cfg == nil || cfg.ConfigPath == "" {
			fmt.Println("No configuration file found.")
			return
		}
		fmt.Println(cfg.ConfigPath)
	},
}

func init() {
	rootCmd.AddCommand(whereCmd)
}
