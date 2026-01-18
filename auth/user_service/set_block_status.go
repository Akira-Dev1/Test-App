package user_service

import (
	"context"
	"fmt"
	"time"

	"auth/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SetUserBlockStatusRequest struct {
	UserId       string     `json:"user_id"`
	IsBlocked    *bool       `json:"is_blocked"`
}


func SetUserBlockStatus(user_id_str string, new_status *bool) error {
	if new_status == nil {
		return fmt.Errorf("is_blocked field is required")
	}

	user_id, err := primitive.ObjectIDFromHex(user_id_str)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": user_id}
	update := bson.M{
		"$set": bson.M{
			"is_blocked": *new_status,
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