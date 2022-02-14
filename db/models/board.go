package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Board struct {
	_ID         primitive.ObjectID `bson:"_id"`
	Description string             `bson:"description"`
	Url         string             `bson:"url"`
	LastPing    int64              `bson:"lastPing"`
	IsUp        bool               `bson:"isUp"`
}
