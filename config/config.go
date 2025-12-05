package config

import "sync/atomic"

type Config struct {
	OpenAIKey    string  `mapstructure:"openai_api_key"`
	Model        string  `mapstructure:"model"`
	Temperature  float32 `mapstructure:"temperature"`
	CustomPrompt string  `mapstructure:"custom_prompt"`
	BasePrompt   string  `mapstructure:"base_prompt"`
	MaxTokens    *int    `mapstructure:"max_tokens"`
}

// configValue holds the current configuration atomically.
var configValue atomic.Value

// Load sets the current configuration.
func store(cfg *Config) {
	configValue.Store(cfg)
}

// Get returns the current configuration.
func Get() *Config {
	v := configValue.Load()
	if v == nil {
		return nil
	}

	return v.(*Config)
}
