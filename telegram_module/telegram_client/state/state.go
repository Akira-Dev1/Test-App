package state

import (
    "sync"
)

var (
    userStates = make(map[int64]string)
    mu         sync.RWMutex
)

func SetUserState(chatID int64, state string) {
    mu.Lock()
    defer mu.Unlock()
    userStates[chatID] = state
}

func GetUserState(chatID int64) string {
    mu.RLock()
    defer mu.RUnlock()
    return userStates[chatID]
}

func ClearUserState(chatID int64) {
    mu.Lock()
    defer mu.Unlock()
    delete(userStates, chatID)
}