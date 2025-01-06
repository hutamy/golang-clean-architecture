package entity

import "time"

type Log struct {
	ID        string    `bson:"_id,omitempty"`
	Message   string    `bson:"message"`
	Level     string    `bson:"level"`
	CreatedAt time.Time `bson:"created_at"`
}
