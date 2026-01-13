package user_service

import (
	"context"
	"time"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"auth/storage"
)


type UserList struct {
	ID        primitive.ObjectID     `bson:"_id,omitempty"`
	Name      string                 `bson:"name"`
}


func GetUserList() ([]UserList, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var users []UserList

	projection := bson.M{
		"_id":  1,
		"name": 1,
	}

	cursor, err := storage.GetUserCollection().Find(ctx, bson.M{}, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user UserList
		if err := cursor.Decode(&user); err != nil {
			log.Printf("Error decoding user basic info: %v", err)
			continue
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}