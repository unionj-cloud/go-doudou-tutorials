package httpsrv

import (
	"net/http"

	ddmodel "github.com/unionj-cloud/go-doudou/framework/http/model"
)

type WordcloudBffHandler interface {
	Upload(w http.ResponseWriter, r *http.Request)
	TaskPage(w http.ResponseWriter, r *http.Request)
}

func Routes(handler WordcloudBffHandler) []ddmodel.Route {
	return []ddmodel.Route{
		{
			"Upload",
			"POST",
			"/upload",
			handler.Upload,
		},
		{
			"TaskPage",
			"POST",
			"/task/page",
			handler.TaskPage,
		},
	}
}
