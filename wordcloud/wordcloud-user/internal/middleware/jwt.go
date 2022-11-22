package middleware

import (
	service "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/transport/httpsrv"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest/httprouter"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/sqlext/wrapper"
	"net/http"
	"strings"
)

func Jwt(db wrapper.GddDB) func(inner http.Handler) http.Handler {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			paramsFromCtx := httprouter.ParamsFromContext(r.Context())
			routeName := paramsFromCtx.MatchedRouteName()
			if !httpsrv.RouteAnnotationStore.HasAnnotation(routeName, "@role") {
				inner.ServeHTTP(w, r)
				return
			}
			authHeader := r.Header.Get("Authorization")
			tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

			userId, valid, err := service.ValidateToken(r.Context(), tokenString, db)
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
