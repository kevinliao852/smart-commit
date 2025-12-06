package generator

import (
	"fmt"

	"github.com/kevinliao852/smart-commit/pkg/git"
	"github.com/kevinliao852/smart-commit/pkg/llm"
)

type Generator struct {
	client llm.Client
}

func New(client llm.Client) *Generator {
	return &Generator{client: client}
}

func (g *Generator) Generate() (string, error) {
	diff, err := git.GetDiff()
	if err != nil {
		return "", fmt.Errorf("failed to read git diff: %w", err)
	}

	if diff == "" {
		return "No staged changes.", nil
	}

	msg, err := g.client.GenerateCommitMessage(diff)
	if err != nil {
		return "", fmt.Errorf("AI generation failed: %w", err)
	}

	return msg, nil
}
