package bd

import (
	"context"
	"time"

	"github.com/nativeron/GoTwitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*recibe email y checkea si existe*/
func CheckUserExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("tw")
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
