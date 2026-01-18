package user_service

import (
	"context"
	"fmt"
	"time"

	"auth/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateUserNameRequest struct {
	UserId       string     `json:"user_id,omitempty"`
	NewName      string     `json:"new_name"`
}


func UpdateUserName(user_id_str string, new_name string) error {
	user_id, err := primitive.ObjectIDFromHex(user_id_str)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": user_id}
	update := bson.M{
		"$set": bson.M{
			"name": new_name,
		},
	}

	result, err := storage.GetUserCollection().UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}