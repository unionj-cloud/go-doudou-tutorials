package httpsrv

import (
	"net/http"

	ddmodel "github.com/unionj-cloud/go-doudou/framework/http/model"
)

type WordcloudSegHandler interface {
	Seg(w http.ResponseWriter, r *http.Request)
}

func Routes(handler WordcloudSegHandler) []ddmodel.Route {
	return []ddmodel.Route{
		{
			"Seg",
			"POST",
			"/seg",
			handler.Seg,
		},
	}
}
