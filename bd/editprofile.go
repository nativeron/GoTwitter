package bd

import (
	"context"
	"time"

	"github.com/nativeron/GoTwitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*modifico perfil de usuario*/
func EditProfile(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("tw")
	col := db.Collection("users")

	register := make(map[string]interface{})
	if len(u.Name) > 0 {
		register["name"] = u.Name
	}
	if len(u.Surname) > 0 {
		register["surname"] = u.Surname
	}
	if len(u.Banner) > 0 {
		register["banner"] = u.Banner
	}
	if len(u.Bio) > 0 {
		register["bio"] = u.Bio
	}
	if len(u.Ubi) > 0 {
		register["ubi"] = u.Ubi
	}
	if len(u.Web) > 0 {
		register["web"] = u.Web
	}
	if len(u.Avatar) > 0 {
		register["avatar"] = u.Avatar
	}
	register["birth"] = u.Birth

	updtString := bson.M{
		"$set": register,
	}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updtString)
	if err != nil {
		return false, err
	}
	return true, nil
}
