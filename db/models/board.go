package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Board struct {
	ID        primitive.ObjectID `bson:"_id"`
	Shortcut  string             `bson:"shortcut"`
	Name      string             `bson:"name"`
	Url       string             `bson:"url"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
	// for the next update
	// LastPing       time.Time              `bson:"lastPing"`
	// IsUp           bool               `bson:"isUp"`
	// PingRetryCount int                `bson:"pingRetryCount"`
}
