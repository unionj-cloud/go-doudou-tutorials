package httpsrv

import (
	"net/http"

	ddmodel "github.com/unionj-cloud/go-doudou/framework/http/model"
)

type AnnotationHandler interface {
	GetGuest(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetAdmin(w http.ResponseWriter, r *http.Request)
}

func Routes(handler AnnotationHandler) []ddmodel.Route {
	return []ddmodel.Route{
		{
			"GetGuest",
			"GET",
			"/guest",
			handler.GetGuest,
		},
		{
			"GetUser",
			"GET",
			"/user",
			handler.GetUser,
		},
		{
			"GetAdmin",
			"GET",
			"/admin",
			handler.GetAdmin,
		},
	}
}

var RouteAnnotationStore = ddmodel.AnnotationStore{
	"GetUser": {
		{
			Name: "@role",
			Params: []string{
				"USER",
				"ADMIN",
			},
		},
	},
	"GetAdmin": {
		{
			Name: "@role",
			Params: []string{
				"ADMIN",
			},
		},
	},
}
