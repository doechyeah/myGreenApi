package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct {
	ID            primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Time_created  primitive.Timestamp `json:"time_created" bson:"time_created"`
	HumidityLevel int                 `json:"humidityLevel" bson:"humidity"`
	Temperature   int                 `json:"temperature" bson:"temperature"`
	UserID        primitive.ObjectID  `json:"user_id" bson:"userid,omitempty"`
}

type User struct {
	ID      primitive.ObjectID `json:"id"`
	Name    string             `json:"name"`
	devices []Device
}
