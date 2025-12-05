package config 
import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Конфиг запуска приложения
type Config struct {
	DataBaseURL string
	GRPCPort 	string
}

func ConfigUP() *Config {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error in config up - godotenv")
	}

	cfg := &Config{
		DataBaseURL: os.Getenv("DB"),
		GRPCPort: os.Getenv("PORT"),
	}

	if cfg.DataBaseURL == "" {
		log.Fatal("database url is bad")
	}

	if cfg.GRPCPort == "" {
		log.Fatal("grpcport is bad")
	}

	return cfg
}