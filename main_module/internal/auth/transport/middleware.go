package transport

import (
	"context"
	"net/http"
	"strings"

	authDomain 	"main_module/internal/auth/domain"
	authInfra 	"main_module/internal/auth/infrastructure"
)

type ctxKey string

const userKey ctxKey = "user"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(header, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "bad auth header", http.StatusUnauthorized)
			return 
		}

		userClaims, err := authInfra.ParseToken(parts[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return 
		}

		user := authDomain.UserContext{
			UserID: userClaims.UserID,
			Permissions: userClaims.Permissions,
			Blocked: userClaims.Blocked,
		}

		ctx := context.WithValue(r.Context(), userKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UserFromContext(ctx context.Context) (authDomain.UserContext, bool) {
	user, ok := ctx.Value(userKey).(authDomain.UserContext)
	return user, ok
}