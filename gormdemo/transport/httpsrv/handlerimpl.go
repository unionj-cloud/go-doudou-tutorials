/**
* Generated by go-doudou v2.0.5.
* You can edit it as your need.
 */
package httpsrv

import (
	"context"
	"encoding/json"
	service "gormdemo"
	"gormdemo/dto"
	"net/http"

	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
)

type GormdemoHandlerImpl struct {
	gormdemo service.Gormdemo
}

func (receiver *GormdemoHandlerImpl) PageUsers(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx   context.Context
		query dto.PageQuery
		data  dto.PageRet
		err   error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&query); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(query); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	data, err = receiver.gormdemo.PageUsers(
		ctx,
		query,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data dto.PageRet `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}

func NewGormdemoHandler(gormdemo service.Gormdemo) GormdemoHandler {
	return &GormdemoHandlerImpl{
		gormdemo,
	}
}
