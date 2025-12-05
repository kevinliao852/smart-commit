package cmd

import (
	"fmt"
	"os"

	"smart-commit/config"

	"github.com/spf13/cobra"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "smart-commit",
	Short: "A CLI tool to generate smart commit messages using AI",
	Long: `smart-commit is a command-line tool that leverages AI to generate
intelligent commit messages based on the changes in your git repository.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load(configFile)
		if err != nil {
			fmt.Println("Error loading config:", err)
		}

		fmt.Fprintln(os.Stderr, "-----------------------------------")
		fmt.Fprintln(os.Stderr, "Configuration loaded successfully.")
		fmt.Fprintf(os.Stderr, "Using model: %s\n", cfg.Model)
		fmt.Fprintf(os.Stderr, "Temperature set to: %.2f\n", cfg.Temperature)
		fmt.Fprintf(os.Stderr, "Max Tokens set to: %d\n", *cfg.MaxTokens)
		fmt.Fprintln(os.Stderr, "-----------------------------------")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(
		&configFile,
		"config",
		"",
		"Path to custom config file (optional)",
	)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
