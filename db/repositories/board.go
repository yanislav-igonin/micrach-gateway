package repositories

import (
	"context"

	Models "micrach-gateway/db/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BoardsRepository interface {
	Create(b *Models.Board) error
	GetAll() (*[]Models.Board, error)
}

type boardsRepository struct {
	collection *mongo.Collection
}

var Boards BoardsRepository

func Init(db *mongo.Database) {
	Boards = &boardsRepository{
		collection: db.Collection("boards"),
	}
}

// Create board record
func (r *boardsRepository) Create(b *Models.Board) error {
	err := r.collection.Database().Client().Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	_, err = r.collection.InsertOne(context.TODO(), b)
	if err != nil {
		return err
	}
	return nil
}

func (r *boardsRepository) GetAll() (*[]Models.Board, error) {
	var boards []Models.Board
	cursor, err := r.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var board Models.Board
		err = cursor.Decode(&board)
		if err != nil {
			return nil, err
		}
		boards = append(boards, board)
	}
	return &boards, nil
}
