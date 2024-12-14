package config

import (
	"gin-web-server/database/mongo"
	"gin-web-server/server/gin"
	"os"
)

// Config struct represents app configuration data
type Config struct {
	Database  mongo.Config
	WebServer gin.Config
}

// CreateConfig : creates configuration struct that contains database connection information, server settings etc.
func CreateConfig() Config {
	mongoCfg := mongo.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     "",
		Password: "",
	}

	wsCfg := gin.Config{
		Port: os.Getenv("WEB_PORT"),
		Host: os.Getenv("WEB_HOST"),
	}

	return Config{
		Database:  mongoCfg,
		WebServer: wsCfg,
	}
}
