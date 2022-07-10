package httpsrv

import (
	"net/http"

	ddmodel "github.com/unionj-cloud/go-doudou/framework/http/model"
)

type TestsvcHandler interface {
	PageUsers(w http.ResponseWriter, r *http.Request)
	LogIn(w http.ResponseWriter, r *http.Request)
	LogIn2(w http.ResponseWriter, r *http.Request)
}

func Routes(handler TestsvcHandler) []ddmodel.Route {
	return []ddmodel.Route{
		{
			"PageUsers",
			"POST",
			"/page/users",
			handler.PageUsers,
		},
		{
			"LogIn",
			"POST",
			"/log/in",
			handler.LogIn,
		},
		{
			"LogIn2",
			"POST",
			"/log/in/2",
			handler.LogIn2,
		},
	}
}

var RouteAnnotationStore = ddmodel.AnnotationStore{}
