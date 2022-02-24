package middleware

import (
	"github.com/gobwas/glob"
	"github.com/jmoiron/sqlx"
	service "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user"
	"net/http"
	"strings"
)

func Jwt(g glob.Glob, conn *sqlx.DB) func(inner http.Handler) http.Handler {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if g.Match(r.URL.Path) {
				inner.ServeHTTP(w, r)
				return
			}
			authHeader := r.Header.Get("Authorization")
			tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

			userId, valid, err := service.ValidateToken(r.Context(), tokenString, conn)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if !valid {
				w.WriteHeader(401)
				w.Write([]byte("Unauthorised.\n"))
				return
			}
			inner.ServeHTTP(w, r.WithContext(service.NewTokenContext(service.NewUserIdContext(r.Context(), userId), tokenString)))
		})
	}
}
