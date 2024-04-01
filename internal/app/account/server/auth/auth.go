package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/cast"
	"net/http"
	"strings"
)

type ctxKey string

const userId ctxKey = "userId"

func GetUserId(ctx context.Context) int {
	return cast.ToInt(ctx.Value(userId))
}

func Middleware(key string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		auths := strings.SplitN(token, " ", 2)
		if len(auths) == 2 && auths[0] == "Bearer" {
			claims := jwt.MapClaims{}
			_, err := jwt.ParseWithClaims(auths[1], &claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(key), nil
			})
			if err == nil {
				if subject, ok := claims["sub"]; ok {
					nextCtx := context.WithValue(r.Context(), userId, subject)
					r = r.WithContext(nextCtx)
				}
			}
		}
		next.ServeHTTP(w, r)
	})
}
