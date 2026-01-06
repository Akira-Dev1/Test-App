package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Email     string             `bson:"email,omitempty"`
	GithubID  *string            `bson:"github_id,omitempty"`
	YandexID  *string            `bson:"yandex_id,omitempty"`
	Role string `bson:"role"`
	CreatedAt time.Time          `bson:"created_at"`
}


// bson - binary json, более экономный тип данных используемый монго
// primitive.ObjectID - базовый тип данных для id в mongo
// omitempty - поле не создается если значения нет