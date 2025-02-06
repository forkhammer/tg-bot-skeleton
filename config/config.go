package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
)

type DBConfig struct {
	Host   string `env:"EXAMPLE_BOT_POSTGRES_HOST"`
	Port   int    `env:"EXAMPLE_BOT_POSTGRES_PORT"`
	DBName string `env:"EXAMPLE_BOT_POSTGRES_DB"`
	User   string `env:"EXAMPLE_BOT_POSTGRES_USER"`
	Pass   string `env:"EXAMPLE_BOT_POSTGRES_PASSWORD"`
}

type Config struct {
	Host         string `env:"EXAMPLE_BOT_HOST"`
	Port         int    `env:"EXAMPLE_BOT_PORT"`
	Token        string `env:"EXAMPLE_BOT_TOKEN"`
	AdminUrl     string `env:"EXAMPLE_BOT_ADMIN_URL"`
	AdminToken   string `env:"EXAMPLE_BOT_ADMIN_TOKEN"`
	Db           DBConfig
	SentryDSN    string `env:"EXAMPLE_BOT_SENTRY_DSN"`
	UserPassword string `env:"EXAMPLE_BOT_USER_PASSWORD"`
}

func NewConfig() *Config {
	config := Config{}
	config.parseConfig()
	return &config
}

func (c *Config) parseConfig() {
	if err := env.Parse(c); err != nil {
		log.Fatal("Can't parse config")
	}
	if err := env.Parse(&c.Db); err != nil {
		log.Fatal("Can't parse db config")
	}
}

func (c *Config) GetHostPort() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

var Settings = NewConfig()
