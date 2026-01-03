package storage

import (
	"sync"
	"time"

	"auth/domain"
)

var (
	codes = make(map[string]domain.CodeState)
	codeMu sync.RWMutex
)

func SaveCode(code domain.CodeState) {
    codeMu.Lock()
    defer codeMu.Unlock()
    codes[code.EntryToken] = code
}

func GetCode(entryToken string) (domain.CodeState, bool) {
    codeMu.RLock()
    defer codeMu.RUnlock()
    state, ok := codes[entryToken]
    return state, ok
}

func DeleteCode(code string) {
	codeMu.Lock()
	defer codeMu.Unlock()
	delete(codes, code)
}

func CleanupExpiredCodes() {
	codeMu.Lock()
	defer codeMu.Unlock()

	now := time.Now()
	for k, v := range codes {
		if v.ExpiresAt.Before(now) {
			delete(codes, k)
		}
	}
}


func AllCodes() map[string]domain.CodeState {
    codeMu.RLock()
    defer codeMu.RUnlock()
    copyMap := make(map[string]domain.CodeState)
    for k, v := range codes {
        copyMap[k] = v
    }
    return copyMap
}
