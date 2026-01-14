package storage

import (
	"context"
	"log"
	"time"
	"errors"

	"auth/config"
	"auth/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"


)

var userCollection *mongo.Collection

func InitUserCollection(dbName, collName string) {
	userCollection = config.Client.Database(dbName).Collection(collName)
	log.Println("User collection initialized")
}

// Найти пользователя по email
func FindUserByEmail(email string) (*domain.User, error) {
	// context — это контейнер с управлением жизненным циклом операции.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user domain.User
	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Создать нового пользователя
func CreateUser(user domain.User) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user.CreatedAt = time.Now()
	user.RefreshTokens = []string{}
	
	res, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	user.ID = res.InsertedID.(primitive.ObjectID)
	return &user, nil
}

// Добавить Refresh Code
func AddRefreshToken(userID primitive.ObjectID, token string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := userCollection.UpdateOne(
        ctx,
        bson.M{"_id": userID},
        bson.M{
            "$push": bson.M{
                "refresh_tokens": token,
            },
        },
    )
    return err
}

func RemoveRefreshToken(userID primitive.ObjectID, token string) error {
	filter := bson.M{
		"_id": userID,
	}

	update := bson.M{
		"$pull": bson.M{
			"refresh_tokens": token,
		},
	}

	res, err := userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}


// GetUserCollection - геттер для коллекции
func GetUserCollection() *mongo.Collection {
    return userCollection
}