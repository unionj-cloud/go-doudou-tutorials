/**
* Generated by go-doudou v2.0.8.
* Don't edit!
 */
package httpsrv

import (
	"net/http"

	"github.com/unionj-cloud/go-doudou/v2/framework"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
)

type GormdemoHandler interface {
	PostGenTUser(w http.ResponseWriter, r *http.Request)
	GetGenTUser_Id(w http.ResponseWriter, r *http.Request)
	PutGenTUser(w http.ResponseWriter, r *http.Request)
	DeleteGenTUser_Id(w http.ResponseWriter, r *http.Request)
	GetGenTUsers(w http.ResponseWriter, r *http.Request)
}

func Routes(handler GormdemoHandler) []rest.Route {
	return []rest.Route{
		{
			Name:        "PostGenTUser",
			Method:      "POST",
			Pattern:     "/gen/t/user",
			HandlerFunc: handler.PostGenTUser,
		},
		{
			Name:        "GetGenTUser_Id",
			Method:      "GET",
			Pattern:     "/gen/t/user/:id",
			HandlerFunc: handler.GetGenTUser_Id,
		},
		{
			Name:        "PutGenTUser",
			Method:      "PUT",
			Pattern:     "/gen/t/user",
			HandlerFunc: handler.PutGenTUser,
		},
		{
			Name:        "DeleteGenTUser_Id",
			Method:      "DELETE",
			Pattern:     "/gen/t/user/:id",
			HandlerFunc: handler.DeleteGenTUser_Id,
		},
		{
			Name:        "GetGenTUsers",
			Method:      "GET",
			Pattern:     "/gen/t/users",
			HandlerFunc: handler.GetGenTUsers,
		},
	}
}

var RouteAnnotationStore = framework.AnnotationStore{}
