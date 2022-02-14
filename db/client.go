package db

import (
	"context"
	"log"
	Config "micrach-gateway/config"
	BoardsRepository "micrach-gateway/db/repositories"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var Database *mongo.Database

func Init() {
	clientOptions := options.Client().ApplyURI(Config.Db.Url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Panicln(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Panicln(err)
	}
	Database = client.Database("micrach-gateway")
	BoardsRepository.Init(Database)
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Panicln(err)
		}
	}()
	log.Println("database - online")
}
