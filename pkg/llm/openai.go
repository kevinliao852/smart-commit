package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type OpenAIClient struct {
	apiKey       string
	model        string
	customPrompt string
	basePrompt   string
	maxTokens    *int
}

func NewOpenAIClient(
	apiKey string,
	model string,
	customPrompt string,
	basePrompt string,
	maxTokens *int,
) *OpenAIClient {
	return &OpenAIClient{
		apiKey:       apiKey,
		model:        model,
		customPrompt: customPrompt,
		basePrompt:   basePrompt,
		maxTokens:    maxTokens,
	}
}

var _ Client = (*OpenAIClient)(nil)

func (c *OpenAIClient) GenerateCommitMessage(diff string) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"

	prompt := c.basePrompt
	if prompt == "" {
		prompt = "Generate a concise, high-quality git commit message for this diff:\n\n%s"
	}
	prompt = prompt + "\n\n" + c.customPrompt

	req := map[string]any{
		"model": c.model,
		"messages": []map[string]string{
			{"role": "user", "content": fmt.Sprintf(prompt, diff)},
		},
		"max_tokens": c.maxTokens,
	}

	body, _ := json.Marshal(req)

	r, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	r.Header.Set("Authorization", "Bearer "+c.apiKey)
	r.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return "", err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Warning: failed to close response body:", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		var errResp struct {
			Error struct {
				Message string `json:"message"`
			} `json:"error"`
		}

		err := json.NewDecoder(resp.Body).Decode(&errResp)
		if err != nil {
			return "", err
		}

		fmt.Println("OpenAI API error message:", errResp.Error.Message)
		return "", fmt.Errorf("OpenAI API returned status code %d", resp.StatusCode)
	}

	var decoded struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&decoded); err != nil {
		return "", err
	}

	if len(decoded.Choices) == 0 {
		return "", fmt.Errorf("no choices returned from OpenAI API")
	}

	return decoded.Choices[0].Message.Content, nil
}
