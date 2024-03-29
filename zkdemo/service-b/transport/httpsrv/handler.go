/**
* Generated by go-doudou v2.0.6.
* Don't edit!
 */
package httpsrv

import (
	"net/http"

	"github.com/unionj-cloud/go-doudou/v2/framework"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
)

type ServiceBHandler interface {
	GetDeptById(w http.ResponseWriter, r *http.Request)
}

func Routes(handler ServiceBHandler) []rest.Route {
	return []rest.Route{
		{
			Name:        "GetDeptById",
			Method:      "GET",
			Pattern:     "/dept/by/id",
			HandlerFunc: handler.GetDeptById,
		},
	}
}

var RouteAnnotationStore = framework.AnnotationStore{}
