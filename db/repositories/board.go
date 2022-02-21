package repositories

import (
	"context"
	"time"

	Models "micrach-gateway/db/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	existedBoard, err := r.GetByShortcut(b.Shortcut)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}
	if existedBoard != nil {
		// TODO: add some meaningful error
		panic("board already exists")
	}

	b.ID, b.CreatedAt, b.UpdatedAt = primitive.NewObjectID(), time.Now(), time.Now()

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

func (r *boardsRepository) GetByShortcut(shortcut string) (*Models.Board, error) {
	var board Models.Board
	result := r.collection.FindOne(context.TODO(), bson.D{{"shortcut", shortcut}})
	if result.Err() != nil {
		if result.Err() != mongo.ErrNoDocuments {
			return nil, result.Err()
		}
	}
	err := result.Decode(&board)
	if err != nil {
		return nil, err
	}
	return &board, nil
}
