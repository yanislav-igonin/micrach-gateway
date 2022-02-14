package db

import (
	"context"
	"log"
	Config "micrach-gateway/config"
	Models "micrach-gateway/db/models"

	// BoardsRepository "micrach-gateway/db/repositories"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var Database *mongo.Database
var BoardsCollection *mongo.Collection

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
	// BoardsRepository.Init(Database)
	BoardsCollection = Database.Collection("boards")
	log.Println("database - online")
}

func Disconnect() {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Panicln(err)
	}
}

func CreateBoard(b *Models.Board) error {
	_, err := BoardsCollection.InsertOne(context.TODO(), b)
	if err != nil {
		return err
	}
	return nil
}
