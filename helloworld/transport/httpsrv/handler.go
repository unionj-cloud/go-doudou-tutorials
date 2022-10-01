/**
* Generated by go-doudou v1.2.4.
* Don't edit!
 */
package httpsrv

import (
	"net/http"

	ddmodel "github.com/unionj-cloud/go-doudou/framework/http/model"
)

type HelloworldHandler interface {
	Greeting(w http.ResponseWriter, r *http.Request)
	Bye(w http.ResponseWriter, r *http.Request)
}

func Routes(handler HelloworldHandler) []ddmodel.Route {
	return []ddmodel.Route{
		{
			"Greeting",
			"POST",
			"/greeting",
			handler.Greeting,
		},
		{
			"Bye",
			"POST",
			"/bye",
			handler.Bye,
		},
	}
}

var RouteAnnotationStore = ddmodel.AnnotationStore{}
