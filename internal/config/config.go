package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/caarlos0/env/v9"
)

type Config struct {
	FilePath  string `env:"PASLOK_FILE_PATH" envDefault:"~/.paslok/.paslok"`
	CipherKey string `env:"PASLOK_CIPHER_KEY"`
}

func (c *Config) ReadConfig(cfg *Config) error {
	if err := env.Parse(cfg); err != nil {
		return fmt.Errorf("parse env: %w", err)
	}

	if strings.HasPrefix(cfg.FilePath, "~/") {
		home, _ := os.UserHomeDir()
		cfg.FilePath = filepath.Join(home, cfg.FilePath[2:])
	}

	return nil
}
