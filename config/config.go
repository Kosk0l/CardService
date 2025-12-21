package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
)

// Конфиг запуска приложения
type Config struct {
	App AppConfig
	DB  DBConfig
}

// Конфиг для миграции
type ConfigMigrator struct {
	DB DBConfig
}

type AppConfig struct {
	Port string `env:"PORT" envDefault:"44044"`
}

type DBConfig struct {
	Host string `env:"DB_HOST" envDefault:"localhost"`
	Port string `env:"DB_PORT" envDefault:"5432"`
	User string `env:"DB_USER,required"`
	Pass string `env:"DB_PASS,required"`
	Name string `env:"DB_NAME,required"`
	SSL  string `env:"DB_SSL" envDefault:"disable"`
}

func Load() *Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	return &cfg
}

func LoadMigrator() *ConfigMigrator {
	cfg := ConfigMigrator{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	return &cfg
}

func (c ConfigMigrator) DsnLoad() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DB.User,
		c.DB.Pass,
		c.DB.Host,
		c.DB.Port,
		c.DB.Name,
		c.DB.SSL,
	)
}