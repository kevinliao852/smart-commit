package config

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/spf13/viper"
)

func Load(configFile string) (*Config, error) {
	v := viper.New()

	if configFile != "" {
		v.SetConfigFile(configFile)
	} else {
		home, _ := os.UserHomeDir()
		v.AddConfigPath(filepath.Join(home, ".config", "smart-commit"))

		v.AddConfigPath(home)

		v.SetConfigName("config")
		v.SetConfigType("yaml")
	}

	v.SetEnvPrefix("AICOMMIT")
	v.AutomaticEnv()

	v.SetDefault("model", "gpt-4o-mini")
	v.SetDefault("temperature", 0.3)
	v.SetDefault("max_tokens", 1024)

	err := autoBindEnv(v, Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to bind env vars: %w", err)
	}

	if err := v.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Info: no config file found, proceeding with defaults + environment.")
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	cfg.ConfigPath = v.ConfigFileUsed()

	store(&cfg)

	return &cfg, nil
}

func autoBindEnv(v *viper.Viper, cfg interface{}) error {
	val := reflect.TypeOf(cfg)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		key := field.Tag.Get("mapstructure")
		if key != "" {
			err := v.BindEnv(key)
			if err != nil {
				return fmt.Errorf("failed to bind env for key %s: %w", key, err)
			}
		}
	}
	return nil
}
