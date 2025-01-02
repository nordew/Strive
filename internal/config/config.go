package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	PostgresURL string `env:"POSTGRES_URL"`
	HTTPPort    int    `env:"HTTP_PORT"`
	BOTToken    string `env:"BOT_TOKEN"`
	WebAppURL   string `env:"WEB_APP_URL"`
}

var (
	config *Config
	once   sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		config = &Config{}
		if err := cleanenv.ReadEnv(config); err != nil {
			log.Fatalf("failed to parse config: %v", err)
		}
	})

	return config
}
