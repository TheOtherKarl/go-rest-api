package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	PORT string
	DB   string
}

func GetConfig() ApiConfig {

	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)

	}

	config := ApiConfig{
		PORT: os.Getenv("PORT"),
		DB:   os.Getenv("DB"),
	}

	return config
}
