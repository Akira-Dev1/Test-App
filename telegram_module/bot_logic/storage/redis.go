package storage

import (
	"context"
	"crypto/rand"
	"strconv"
	"time"
	"log"
)

func CreateUser(chat_id int) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    chatIDStr := strconv.Itoa(chat_id)
    entry_token := rand.Text()
    
    log.Printf("CreateUser: creating chat_id=%s with token=%s", chatIDStr, entry_token)
    err := AuthRedis.HSet(ctx, chatIDStr, 
        "entry_token", entry_token, 
        "status", "anonymous",
    ).Err() 
    
    if err != nil {
        log.Printf("CreateUser: error saving user data: %v", err)
        return err
    }

    err = AuthRedis.HSet(ctx, "ids", chatIDStr, chatIDStr).Err()
    if err != nil {
        log.Printf("CreateUser: error adding to ids list: %v", err)
        return err
    }
    
    log.Printf("CreateUser: successfully created user %s", chatIDStr)
    return nil
}

func UpdateUser(chat_id int) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    chatIDStr := strconv.Itoa(chat_id)
    entry_token := rand.Text()
    
    log.Printf("UpdateUser: updating chat_id=%s with token=%s", chatIDStr, entry_token)

    err := AuthRedis.HSet(ctx, chatIDStr, "entry_token", entry_token).Err()
    if err != nil {
        log.Printf("UpdateUser: error updating entry_token: %v", err)
        return err
    }

    err = AuthRedis.HSet(ctx, "ids", chatIDStr, chatIDStr).Err()
    if err != nil {
        log.Printf("UpdateUser: error updating ids list: %v", err)
        return err
    }
    
    log.Printf("UpdateUser: successfully updated user %s", chatIDStr)
    return nil
}

func GetEntryToken(chat_id int) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	entry_token, err := AuthRedis.HGet(ctx, strconv.Itoa(chat_id), "entry_token").Result()
	if err != nil {
		return "", err
	}

	return entry_token, nil
}

func GetRefreshToken(chat_id int) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	refresh_token, err := AuthRedis.HGet(ctx, strconv.Itoa(chat_id), "refresh_token").Result()
	if err != nil {
		return "", err
	}

	return refresh_token, nil
}