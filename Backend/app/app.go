package app

import (
	"context"
	"gin-web-server/config"
	"gin-web-server/database/mongo"
	"log"
)

// Run ...
func Run(cfg config.Config) error {

	c := mongo.Config(cfg.Database)
	db, err := mongo.NewMongoDatabase(context.TODO(), c)
	if err != nil {
		return err
	}

	exercise, err := db.GetExerciseByID(context.TODO(), 0)
	if err != nil {
		return err
	}
	log.Println(exercise)
	return nil
}
