package httpsrv

import (
	"net/http"

	ddmodel "github.com/unionj-cloud/go-doudou/framework/http/model"
)

type UsersvcHandler interface {
	PageUsers(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetMe(w http.ResponseWriter, r *http.Request)
	PublicSignUp(w http.ResponseWriter, r *http.Request)
	PublicLogIn(w http.ResponseWriter, r *http.Request)
	UploadAvatar(w http.ResponseWriter, r *http.Request)
	GetPublicDownloadAvatar(w http.ResponseWriter, r *http.Request)
	GetLogout(w http.ResponseWriter, r *http.Request)
	PublicTokenValidate(w http.ResponseWriter, r *http.Request)
}

func Routes(handler UsersvcHandler) []ddmodel.Route {
	return []ddmodel.Route{
		{
			"PageUsers",
			"POST",
			"/page/users",
			handler.PageUsers,
		},
		{
			"User",
			"GET",
			"/user",
			handler.GetUser,
		},
		{
			"Me",
			"GET",
			"/me",
			handler.GetMe,
		},
		{
			"PublicSignUp",
			"POST",
			"/public/sign/up",
			handler.PublicSignUp,
		},
		{
			"PublicLogIn",
			"POST",
			"/public/log/in",
			handler.PublicLogIn,
		},
		{
			"UploadAvatar",
			"POST",
			"/upload/avatar",
			handler.UploadAvatar,
		},
		{
			"PublicDownloadAvatar",
			"GET",
			"/public/download/avatar",
			handler.GetPublicDownloadAvatar,
		},
		{
			"Logout",
			"GET",
			"/logout",
			handler.GetLogout,
		},
		{
			"PublicTokenValidate",
			"POST",
			"/public/token/validate",
			handler.PublicTokenValidate,
		},
	}
}
