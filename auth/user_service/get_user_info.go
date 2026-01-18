package user_service

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"auth/storage"
)

type UserInf struct {
	Name  string `bson:"name"`
}

type UserInfRequest struct {
	UserId string `json:"user_id"`
}

func GetUserInfo(user_id_str string) (UserInf, error) {
	var user UserInf
	
	user_id, err := primitive.ObjectIDFromHex(user_id_str)
	if err != nil {
		return user, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	// Создаем options для FindOne
	findOptions := options.FindOne()
	findOptions.SetProjection(bson.M{
		"name": 1,
		"_id":  0, // Явно исключаем _id из результатов
	})

	err = storage.GetUserCollection().FindOne(
		ctx, 
		bson.M{"_id": user_id}, 
		findOptions, // Передаем options.FindOneOptions
	).Decode(&user)
	
	if err != nil {
		return user, err
	}
	return user, nil
}