package storage

import (
	"sync"
	"time"

	"auth/domain"
)

var (
	authStates = make(map[string]domain.AuthState)
	authMu sync.RWMutex
)

func SaveAuthState(state domain.AuthState) {
	authMu.Lock()
	defer authMu.Unlock()
	authStates[state.EntryToken] = state
}

func GetAuthState(entryToken string) (domain.AuthState, bool) {
	authMu.RLock()
	defer authMu.RUnlock()
	state, ok := authStates[entryToken]
	return state, ok
}

func DeleteAuthState(entryToken string) {
	authMu.Lock()
	defer authMu.Unlock()
	delete(authStates, entryToken)
}

func CleanupExpiredAuthStates() {
	authMu.Lock()
	defer authMu.Unlock()

	now := time.Now()
	for k, v := range authStates {
		if v.ExpiresAt.Before(now) {
			delete(authStates, k)
		}
	}
}
