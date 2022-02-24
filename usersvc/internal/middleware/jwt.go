package middleware

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gobwas/glob"
	"net/http"
	"os"
	"strings"
)

type ctxKey int

const userIdKey ctxKey = ctxKey(0)

func NewContext(ctx context.Context, id int) context.Context {
	return context.WithValue(ctx, userIdKey, id)
}

func FromContext(ctx context.Context) (int, bool) {
	userId, ok := ctx.Value(userIdKey).(int)
	return userId, ok
}

func Jwt(g glob.Glob) func(inner http.Handler) http.Handler {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if g.Match(r.URL.Path) {
				inner.ServeHTTP(w, r)
				return
			}
			authHeader := r.Header.Get("Authorization")
			tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("BIZ_JWT_SECRET")), nil
			})
			if err != nil || !token.Valid {
				w.WriteHeader(401)
				w.Write([]byte("Unauthorised.\n"))
				return
			}

			claims := token.Claims.(jwt.MapClaims)
			if userId, exists := claims["userId"]; !exists {
				w.WriteHeader(401)
				w.Write([]byte("Unauthorised.\n"))
				return
			} else {
				inner.ServeHTTP(w, r.WithContext(NewContext(r.Context(), int(userId.(float64)))))
			}
		})
	}
}
