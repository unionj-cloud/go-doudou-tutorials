package httpsrv

import (
	"annotation/vo"
	"github.com/gorilla/mux"
	"github.com/unionj-cloud/go-doudou/toolkit/sliceutils"
	"net/http"
)

func Auth(userStore vo.UserStore) func(inner http.Handler) http.Handler {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentRoute := mux.CurrentRoute(r)
			if currentRoute == nil {
				inner.ServeHTTP(w, r)
				return
			}
			routeName := currentRoute.GetName()
			if !RouteAnnotationStore.HasAnnotation(routeName, "@role") {
				inner.ServeHTTP(w, r)
				return
			}
			user, pass, ok := r.BasicAuth()
			if !ok {
				w.Header().Set("WWW-Authenticate", `Basic realm="Provide user name and password"`)
				w.WriteHeader(401)
				w.Write([]byte("Unauthorised.\n"))
				return
			}
			role, exists := userStore[vo.Auth{user, pass}]
			if !exists {
				w.Header().Set("WWW-Authenticate", `Basic realm="Provide user name and password"`)
				w.WriteHeader(401)
				w.Write([]byte("Unauthorised.\n"))
				return
			}
			params := RouteAnnotationStore.GetParams(routeName, "@role")
			if !sliceutils.StringContains(params, role.StringGetter()) {
				w.WriteHeader(403)
				w.Write([]byte("Access denied\n"))
				return
			}
			inner.ServeHTTP(w, r)
		})
	}
}
