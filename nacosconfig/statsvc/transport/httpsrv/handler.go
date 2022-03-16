package httpsrv

import (
	"net/http"

	ddmodel "github.com/unionj-cloud/go-doudou/framework/http/model"
)

type StatsvcHandler interface {
	Add(w http.ResponseWriter, r *http.Request)
	GetEcho(w http.ResponseWriter, r *http.Request)
}

func Routes(handler StatsvcHandler) []ddmodel.Route {
	return []ddmodel.Route{
		{
			"Add",
			"POST",
			"/add",
			handler.Add,
		},
		{
			"Echo",
			"GET",
			"/echo",
			handler.GetEcho,
		},
	}
}
