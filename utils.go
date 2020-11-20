package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config struct
type Config struct {
	Debug bool   `default:"false"`
	Host  string `default:"localhost"`
	Port  string `default:"5432"`
}

// LoadConfig laod all the .env file to read the enviroment config
func LoadConfig() *Config {
	_ = godotenv.Load()

	var cfg Config
	err := envconfig.Process("BACK", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &cfg
}
