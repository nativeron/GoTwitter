package models

type Rel struct {
	UserID    string `bson:"userid" json:"userId"`
	UserRelID string `bson:"userrelid" json:"userRelId"`
}
