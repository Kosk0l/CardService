package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

// Конфиг запуска приложения
type Config struct {
	App AppConfig
	DB  DBConfig
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
}

func Load() *Config {

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	return &cfg
}