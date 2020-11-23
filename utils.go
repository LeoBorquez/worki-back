package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config struct
type Config struct {
	Debug   bool `default:"false"`
	Cors    string
	Host    string
	Port    int
	User    string
	Name    string
	SSLmode string
	Pass    string
}

// LoadConfig load all the .env file to read the enviroment config
func LoadConfig() *Config {
	_ = godotenv.Load()

	var cfg Config
	err := envconfig.Process("BACK", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &cfg
}
