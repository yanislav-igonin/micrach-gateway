package repositories

import (
	"context"

	Db "micrach-gateway/db"
	Models "micrach-gateway/db/models"

	"go.mongodb.org/mongo-driver/bson"
)

// Create board record
func Create(b *Models.Board) error {
	_, err := Db.Database.Collection("boards").InsertOne(context.TODO(), b)
	if err != nil {
		return err
	}
	return nil
}

func GetAll() (*[]Models.Board, error) {
	var boards []Models.Board
	cursor, err := Db.Database.Collection("boards").Find(context.TODO(), bson.D{})
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
