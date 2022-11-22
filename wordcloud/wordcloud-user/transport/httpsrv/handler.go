/**
* Generated by go-doudou v2.0.1.
* Don't edit!
 */
package httpsrv

import (
	"net/http"

	"github.com/unionj-cloud/go-doudou/v2/framework"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
)

type UsersvcHandler interface {
	GetUser(w http.ResponseWriter, r *http.Request)
	GetMe(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	LogIn(w http.ResponseWriter, r *http.Request)
	GetLogout(w http.ResponseWriter, r *http.Request)
	GetTokenValidate(w http.ResponseWriter, r *http.Request)
}

func Routes(handler UsersvcHandler) []rest.Route {
	return []rest.Route{
		{
			Name:        "GetUser",
			Method:      "GET",
			Pattern:     "/user",
			HandlerFunc: handler.GetUser,
		},
		{
			Name:        "GetMe",
			Method:      "GET",
			Pattern:     "/me",
			HandlerFunc: handler.GetMe,
		},
		{
			Name:        "SignUp",
			Method:      "POST",
			Pattern:     "/sign/up",
			HandlerFunc: handler.SignUp,
		},
		{
			Name:        "LogIn",
			Method:      "POST",
			Pattern:     "/log/in",
			HandlerFunc: handler.LogIn,
		},
		{
			Name:        "GetLogout",
			Method:      "GET",
			Pattern:     "/logout",
			HandlerFunc: handler.GetLogout,
		},
		{
			Name:        "GetTokenValidate",
			Method:      "GET",
			Pattern:     "/token/validate",
			HandlerFunc: handler.GetTokenValidate,
		},
	}
}

var RouteAnnotationStore = framework.AnnotationStore{
	"GetUser": {
		{
			Name: "@role",
			Params: []string{
				"USER",
			},
		},
	},
	"GetMe": {
		{
			Name: "@role",
			Params: []string{
				"USER",
			},
		},
	},
	"GetLogout": {
		{
			Name: "@role",
			Params: []string{
				"USER",
			},
		},
	},
}
