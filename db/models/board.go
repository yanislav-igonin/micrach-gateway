package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Board struct {
	_ID         primitive.ObjectID `bson:"_id"`
	ID          string             `bson:"id"`
	Description string             `bson:"description"`
	Url         string             `bson:"url"`
	lastPing    int64              `bson:"lastPing"`
	isUp        bool               `bson:"isUp"`
}
