package user_service

import (
	"context"
	"fmt"
	"time"

	"auth/storage"
	"auth/permissions"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateUserRolesRequest struct {
	UserId       string     `json:"user_id,omitempty"`
	Roles        []string   `json:"roles"`
}


func UpdateUserRoles(user_id_str string, new_roles []string) error {
	user_id, err := primitive.ObjectIDFromHex(user_id_str)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": user_id}
	update := bson.M{
		"$set": bson.M{
			"roles": new_roles,
			"permissions": permissions.ResolvePermissions(new_roles),
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