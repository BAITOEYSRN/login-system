package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB_URL      string `split_words:"DB_URL"`
	DB_HOST     string `split_words:"DB_HOST"`
	DB_PORT     string `split_words:"DB_PORT"`
	DB_USER     string `split_words:"DB_USER"`
	DB_PASSWORD string `split_words:"DB_PASSWORD"`
	DB_NAME     string `split_words:"DB_NAME"`
	DB_SSLMODE  string `split_words:"DB_SSLMODE"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("error processing env config: %w", err)
	}

	return &cfg, nil
}
