package httpsrv

import (
	"net/http"

	ddmodel "github.com/unionj-cloud/go-doudou/framework/http/model"
)

type WordcloudTaskHandler interface {
	TaskSave(w http.ResponseWriter, r *http.Request)
	TaskSuccess(w http.ResponseWriter, r *http.Request)
	TaskFail(w http.ResponseWriter, r *http.Request)
	TaskPage(w http.ResponseWriter, r *http.Request)
}

func Routes(handler WordcloudTaskHandler) []ddmodel.Route {
	return []ddmodel.Route{
		{
			"TaskSave",
			"POST",
			"/task/save",
			handler.TaskSave,
		},
		{
			"TaskSuccess",
			"POST",
			"/task/success",
			handler.TaskSuccess,
		},
		{
			"TaskFail",
			"POST",
			"/task/fail",
			handler.TaskFail,
		},
		{
			"TaskPage",
			"POST",
			"/task/page",
			handler.TaskPage,
		},
	}
}
