package httpsrv

import (
	"net/http"

	ddmodel "github.com/unionj-cloud/go-doudou/framework/http/model"
)

type EnumDemoHandler interface {
	GetKeyboard(w http.ResponseWriter, r *http.Request)
	GetKeyboard2(w http.ResponseWriter, r *http.Request)
	GetKeyboards(w http.ResponseWriter, r *http.Request)
	GetKeyboards2(w http.ResponseWriter, r *http.Request)
	GetKeyboards5(w http.ResponseWriter, r *http.Request)
	Keyboard(w http.ResponseWriter, r *http.Request)
}

func Routes(handler EnumDemoHandler) []ddmodel.Route {
	return []ddmodel.Route{
		{
			"Keyboard",
			"GET",
			"/keyboard",
			handler.GetKeyboard,
		},
		{
			"Keyboard2",
			"GET",
			"/keyboard/2",
			handler.GetKeyboard2,
		},
		{
			"Keyboards",
			"GET",
			"/keyboards",
			handler.GetKeyboards,
		},
		{
			"Keyboards2",
			"GET",
			"/keyboards/2",
			handler.GetKeyboards2,
		},
		{
			"Keyboards5",
			"GET",
			"/keyboards/5",
			handler.GetKeyboards5,
		},
		{
			"Keyboard",
			"POST",
			"/keyboard",
			handler.Keyboard,
		},
	}
}
