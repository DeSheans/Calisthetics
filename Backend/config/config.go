package config

import (
	"gin-web-server/database/mongo"
	"os"
)

// Config struct represents app configuration data
type Config struct {
	Database  mongo.Config
	WebServer WebServer
}

// WebServer struct represents configuration settings that are used to work with web-server
type WebServer struct {
	Port string
}

// CreateConfig : creates configuration struct that contains database connection information, server settings etc.
func CreateConfig() Config {
	mongoCfg := mongo.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     "",
		Password: "",
	}

	wsCfg := WebServer{
		Port: ":8080",
	}

	return Config{
		Database:  mongoCfg,
		WebServer: wsCfg,
	}
}
