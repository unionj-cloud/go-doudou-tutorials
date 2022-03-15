package httpsrv

import (
	"net/http"

	ddmodel "github.com/unionj-cloud/go-doudou/framework/http/model"
)

type WordcloudMakerHandler interface {
	Make(w http.ResponseWriter, r *http.Request)
}

func Routes(handler WordcloudMakerHandler) []ddmodel.Route {
	return []ddmodel.Route{
		{
			"Make",
			"POST",
			"/make",
			handler.Make,
		},
	}
}
