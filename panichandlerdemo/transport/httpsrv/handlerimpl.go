/**
* Generated by go-doudou v2.0.5.
* You can edit it as your need.
 */
package httpsrv

import (
	"context"
	"encoding/json"
	"net/http"
	service "testsvc"
	"testsvc/dto"

	"github.com/pkg/errors"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
)

type TestsvcHandlerImpl struct {
	testsvc service.Testsvc
}

func (receiver *TestsvcHandlerImpl) GetBookNotFoundException(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx context.Context
		re  error
	)
	ctx = _req.Context()
	re = receiver.testsvc.GetBookNotFoundException(
		ctx,
	)
	if re != nil {
		panic(re)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
	}{}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestsvcHandlerImpl) GetConversionFailedException(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx context.Context
		re  error
	)
	ctx = _req.Context()
	re = receiver.testsvc.GetConversionFailedException(
		ctx,
	)
	if re != nil {
		panic(re)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
	}{}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestsvcHandlerImpl) GetBookPage(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx         context.Context
		name        string
		author      string
		pageWrapper struct {
			Page dto.Page `form:"page"`
		}
		re error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	if _, exists := _req.Form["name"]; exists {
		name = _req.FormValue("name")
		if _err := rest.ValidateVar(name, "", "name"); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	} else {
		rest.HandleBadRequestErr(errors.New("missing parameter name"))
	}
	if _, exists := _req.Form["author"]; exists {
		author = _req.FormValue("author")
		if _err := rest.ValidateVar(author, "", "author"); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	} else {
		rest.HandleBadRequestErr(errors.New("missing parameter author"))
	}
	if _err := rest.DecodeForm(&pageWrapper, _req.Form); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(pageWrapper.Page); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	re = receiver.testsvc.GetBookPage(
		ctx,
		name,
		author,
		pageWrapper.Page,
	)
	if re != nil {
		panic(re)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
	}{}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestsvcHandlerImpl) PostBookPage(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx    context.Context
		name   string
		author string
		re     error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	if _, exists := _req.Form["name"]; exists {
		name = _req.FormValue("name")
		if _err := rest.ValidateVar(name, "", "name"); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	} else {
		rest.HandleBadRequestErr(errors.New("missing parameter name"))
	}
	if _, exists := _req.Form["author"]; exists {
		author = _req.FormValue("author")
		if _err := rest.ValidateVar(author, "", "author"); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	} else {
		rest.HandleBadRequestErr(errors.New("missing parameter author"))
	}
	re = receiver.testsvc.PostBookPage(
		ctx,
		name,
		author,
	)
	if re != nil {
		panic(re)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
	}{}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}

func NewTestsvcHandler(testsvc service.Testsvc) TestsvcHandler {
	return &TestsvcHandlerImpl{
		testsvc,
	}
}
