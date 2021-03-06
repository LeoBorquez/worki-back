package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config struct
type Config struct {
	Dev     bool `default:"false"`
	Cors    string
	HostDB  string
	PortDB  string `default:"5432"`
	UserDB  string
	NameDB  string
	PassDB  string
	URI     string `default:""`
	SSLmode string
	Port    string `default:"1323"`
}

// LoadConfig load all the .env file to read the enviroment config
func LoadConfig() *Config {
	// Load the .env file
	env := godotenv.Load()
	if env != nil {
		log.Fatalf("Error loading .env file")
	}

	var cfg Config
	err := envconfig.Process("back", &cfg)
	if err != nil {
		log.Fatalf(err.Error(), log.Ldate)
	}
	return &cfg
}
