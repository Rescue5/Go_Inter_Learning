package configs

import (
	"HttpServer/extra"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(extra.WrapError("Error Loading .env file: ", err))
	}

	return &Config{
		Db:   DbConfig{Dsn: os.Getenv("DSN")},
		Auth: AuthConfig{Secret: os.Getenv("TOKEN")},
	}

}
