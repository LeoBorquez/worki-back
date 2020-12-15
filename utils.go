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
	Port    string `default:"5432"`
	UserDB  string
	NameDB  string
	PassDB  string
	SSLmode string
}

// LoadConfig load all the .env file to read the enviroment config
func LoadConfig() *Config {
	// Load the .env file
	env := godotenv.Load()
	if env != nil {
		log.Fatalf("Error loading .env file")
	}

	var cfg Config
	err := envconfig.Process("BACK", &cfg)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return &cfg
}
