package db

import (
	"context"
	"log"
	Config "micrach-gateway/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Init() {
	Client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(Config.Db.Url))
	if err != nil {
		log.Panicln(err)
	}
	defer func() {
		if err := Client.Disconnect(context.TODO()); err != nil {
			log.Panicln(err)
		}
	}()
	log.Println("database - online")
}
