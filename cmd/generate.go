package cmd

import (
	"fmt"

	"github.com/kevinliao852/smart-commit/config"
	"github.com/kevinliao852/smart-commit/pkg/generator"
	"github.com/kevinliao852/smart-commit/pkg/llm"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a commit message using AI based on staged changes",
	Long: `Generate a commit message using AI based on staged changes.
This command analyzes the staged changes in your git repository
and generates a suitable commit message using an AI language model.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.Get()

		g := generator.New(llm.NewOpenAIClient(
			cfg.OpenAIKey, cfg.Model, cfg.CustomPrompt, cfg.BasePrompt, cfg.MaxTokens,
		))
		msg, err := g.Generate()
		if err != nil {
			return err
		}

		fmt.Println(msg)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
