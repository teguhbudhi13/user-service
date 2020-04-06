package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DBConfig (Database Configuration)
type DBConfig struct {
	Dialect  string
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	SslMode  string
}

// ServerConfig for server configuration
type ServerConfig struct {
	Port string
}

// Config Constructor
type Config struct {
	DB     *DBConfig
	Server *ServerConfig
}

// GetConfig Generate Configuration
func GetConfig() *Config {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			SslMode:  os.Getenv("DB_SSL_MODE"),
		},
		Server: &ServerConfig{
			Port: os.Getenv("SERVER_PORT"),
		},
	}
}
