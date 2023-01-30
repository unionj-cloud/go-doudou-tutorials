/**
* Generated by go-doudou v2.0.5.
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
	GetShelves_ShelfBooks_Book(w http.ResponseWriter, r *http.Request)
	GetBook(w http.ResponseWriter, r *http.Request)
}

func Routes(handler GoStatsHandler) []rest.Route {
	return []rest.Route{
		{
			Name:        "LargestRemainder",
			Method:      "POST",
			Pattern:     "/largest/remainder",
			HandlerFunc: handler.LargestRemainder,
		},
		{
			Name:        "GetShelves_ShelfBooks_Book",
			Method:      "GET",
			Pattern:     "/shelves/:shelf/books/:book",
			HandlerFunc: handler.GetShelves_ShelfBooks_Book,
		},
		{
			Name:        "GetBook",
			Method:      "GET",
			Pattern:     "/book",
			HandlerFunc: handler.GetBook,
		},
	}
}

var RouteAnnotationStore = framework.AnnotationStore{}
