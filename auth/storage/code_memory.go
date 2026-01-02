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
	codes[code.Code] = code
}

func GetCode(code string) (domain.CodeState, bool) {
	codeMu.RLock()
	defer codeMu.RUnlock()
	state, ok := codes[code]
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