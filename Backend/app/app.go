package app

import (
	"context"
	"gin-web-server/app/handlers"
	"gin-web-server/config"
	"gin-web-server/database/mongo"
	"gin-web-server/server/gin"
)

// Run ...
func Run(cfg config.Config) error {

	db, err := mongo.NewMongoDatabase(context.TODO(), cfg.Database)
	if err != nil {
		return err
	}

	hand := handlers.NewConcreteHandler(db)

	server := gin.NewGinServer(cfg.WebServer, hand)
	if err = server.Run(); err != nil {
		return nil
	}

	return nil
}
