/**
* Generated by go-doudou v2.0.3.
* Don't edit!
 */
package httpsrv

import (
	"net/http"

	"github.com/unionj-cloud/go-doudou/v2/framework"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
)

type GoStatsHandler interface {
	LargestRemainder(w http.ResponseWriter, r *http.Request)
}

func Routes(handler GoStatsHandler) []rest.Route {
	return []rest.Route{
		{
			Name:        "LargestRemainder",
			Method:      "POST",
			Pattern:     "/largest/remainder",
			HandlerFunc: handler.LargestRemainder,
		},
	}
}

var RouteAnnotationStore = framework.AnnotationStore{}
