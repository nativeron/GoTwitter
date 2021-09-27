package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*estructura con la que se devuleven los tweets*/
type ReturnTweets struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID  string             `bson:"userid" json:"userId,omitempty"`
	Message string             `bson:"message" json:"message,omitempty"`
	Date    time.Time          `bson:"date" json:"date,omitempty"`
}
