package httpsrv

import (
	service "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff"
	userclient "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/client"
	"github.com/unionj-cloud/go-doudou/framework/ratelimit/redisrate"
	"net/http"
	"strconv"
	"strings"
)

func Auth(userClient *userclient.UsersvcClientProxy) func(inner http.Handler) http.Handler {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
			userVo, _ := userClient.PublicTokenValidate(r.Context(), tokenString)
			if userVo.Id == 0 {
				w.WriteHeader(401)
				w.Write([]byte("Unauthorised.\n"))
				return
			}
			inner.ServeHTTP(w, r.WithContext(service.NewUserIdContext(r.Context(), userVo.Id)))
		})
	}
}

// RedisRateLimit limit rate based on redis
func RedisRateLimit(rdb redisrate.Rediser, fn redisrate.LimitFn) func(inner http.Handler) http.Handler {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userId, _ := service.UserIdFromContext(r.Context())
			limiter := redisrate.NewGcraLimiterLimitFn(rdb, strconv.Itoa(userId), fn)
			if !limiter.Allow() {
				http.Error(w, "too many requests", http.StatusTooManyRequests)
				return
			}
			inner.ServeHTTP(w, r)
		})
	}
}
