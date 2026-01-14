package user_service

import (
	"context"
	"log"
	"time"

	"auth/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RolesResponse struct {
	Roles []string `bson:"roles"`
}

type GetRolesRequest struct {
	UserId string `json:"user_id"`
}

func GetUserRoles(user_id_str string) (roles RolesResponse, err error) {
	user_id, err := primitive.ObjectIDFromHex(user_id_str)
	// if err != nil {
	// 	log.Println("1")
	// 	return , err
	// }

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	// Создаем options для FindOne
	findOptions := options.FindOne()
	findOptions.SetProjection(bson.M{
		"roles": 1,
		"_id":  0, // Явно исключаем _id из результатов
	})

	var roles_col RolesResponse
	err = storage.GetUserCollection().FindOne(
		ctx, 
		bson.M{"_id": user_id}, 
		findOptions, // Передаем options.FindOneOptions
	).Decode(&roles_col)
	if err != nil {
		log.Println("2")
		return roles_col, err
	}
	return roles_col, nil
	
}