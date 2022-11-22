package httpsrv

import (
	service "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff"
	userclient "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/client"
	"net/http"
	"strings"
)

func Auth(userClient userclient.IUsersvcClient) func(inner http.Handler) http.Handler {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
			_, userVo, _ := userClient.PublicTokenValidate(r.Context(), nil, tokenString)
			if userVo.Id == 0 {
				w.WriteHeader(401)
				w.Write([]byte("Unauthorised.\n"))
				return
			}
			inner.ServeHTTP(w, r.WithContext(service.NewUserIdContext(service.NewTokenContext(r.Context(), tokenString), userVo.Id)))
		})
	}
}
