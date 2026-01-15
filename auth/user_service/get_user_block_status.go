package user_service


import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"auth/storage"
)

type UserGetBlockStatus struct {
	IsBlocked bool `bson:"is_blocked"`
}

type UserGetBlockStatusRequest struct {
	UserId string `json:"user_id"`
}

func GetUserBlockStatus(user_id_str string) (block_status UserGetBlockStatus, err error) {
	var status UserGetBlockStatus
	
	user_id, _ := primitive.ObjectIDFromHex(user_id_str)
	if err != nil {
		return status, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	findOptions := options.FindOne()
	findOptions.SetProjection(bson.M{
		"is_blocked": 1,
		"_id":  0,
	})

	err = storage.GetUserCollection().FindOne(
		ctx, 
		bson.M{"_id": user_id}, 
		findOptions, 
	).Decode(&status)
	
	if err != nil {
		return status, err
	}
	return status, nil
}