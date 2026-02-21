package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Rules struct {
		UpperCaseCheck bool `yaml:"uppercase-check"`
		CyrillicCheck bool `yaml:"cyrillic-check"`
		CriticalInfoCheck bool `yaml:"critical-info-check"`
		SpecialSymbolsCheck bool `yaml:"special-symbols-check"`
	}
	DangerousWords []string `yaml:"dangerous-words"`
	AllowedMethods []string `yaml:"allowed-methods"`
	PreventedMethods []string `yaml:"prevented-methods"`
}

func New() (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig("./config/config.yaml", &cfg); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}
	return &cfg, nil
}