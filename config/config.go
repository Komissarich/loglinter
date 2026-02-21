package config

import (
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
		return &Config{
			Rules: struct {
				UpperCaseCheck      bool `yaml:"uppercase-check"`
				CyrillicCheck       bool `yaml:"cyrillic-check"`
				CriticalInfoCheck   bool `yaml:"critical-info-check"`
				SpecialSymbolsCheck bool `yaml:"special-symbols-check"`
			}{
				UpperCaseCheck:      true,
				CyrillicCheck:       true,
				CriticalInfoCheck:   true,
				SpecialSymbolsCheck: true,
			},
			DangerousWords: []string{
				"password", "passwd", "pwd", "token", "api_key", "apikey", "secret",
				"access_token", "refresh_token", "private_key", "auth", "session_id",
			},
			AllowedMethods: []string{
				"Info", "Warn", "Error", "Println", "Print", "InfoContext", "WarnContext",
				"ErrorContext", "DebugContext", "Log", "LogAttrs",
			},
			PreventedMethods: []string{
				// если хочешь запретить какие-то методы — добавь сюда
			},
		}, nil
		// return nil, fmt.Errorf("failed to read config: %w", err) for debugging purposes
	}
	return &cfg, nil
}