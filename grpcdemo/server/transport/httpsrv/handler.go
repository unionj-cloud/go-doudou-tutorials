/**
* Generated by go-doudou v1.3.7.
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
	BiStream(w http.ResponseWriter, r *http.Request)
	ClientStream(w http.ResponseWriter, r *http.Request)
	ServerStream(w http.ResponseWriter, r *http.Request)
}

func Routes(handler HelloworldHandler) []ddmodel.Route {
	return []ddmodel.Route{
		{
			Name:        "Greeting",
			Method:      "POST",
			Pattern:     "/greeting",
			HandlerFunc: handler.Greeting,
		},
		{
			Name:        "Bye",
			Method:      "POST",
			Pattern:     "/bye",
			HandlerFunc: handler.Bye,
		},
		{
			Name:        "BiStream",
			Method:      "POST",
			Pattern:     "/bi/stream",
			HandlerFunc: handler.BiStream,
		},
		{
			Name:        "ClientStream",
			Method:      "POST",
			Pattern:     "/client/stream",
			HandlerFunc: handler.ClientStream,
		},
		{
			Name:        "ServerStream",
			Method:      "POST",
			Pattern:     "/server/stream",
			HandlerFunc: handler.ServerStream,
		},
	}
}

var RouteAnnotationStore = ddmodel.AnnotationStore{}
