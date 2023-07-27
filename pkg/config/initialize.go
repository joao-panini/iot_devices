package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadConfigurations() (Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		err := godotenv.Load("../.env")
		if err != nil {
			return Config{}, err
		}
	}

	return Config{
		Database: DatabaseConfigurations{
			User:           os.Getenv("DB_USER"),
			Password:       os.Getenv("DB_PASSWORD"),
			Host:           os.Getenv("DB_HOST"),
			Port:           os.Getenv("DB_PORT"),
			Name:           os.Getenv("DB_NAME"),
			SSLMode:        os.Getenv("DB_SSL"),
			MigrationsPath: os.Getenv("DB_MIG"),
		},
	}, nil
}
