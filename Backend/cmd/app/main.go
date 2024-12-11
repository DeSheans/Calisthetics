package main

import (
	"gin-web-server/app"
	"gin-web-server/config"
)

func main() {
	cfg := config.CreateConfig()
	if err := app.Run(cfg); err != nil {
		panic(err)
	}
}
