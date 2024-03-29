/**
* Generated by go-doudou v2.1.4.
* You can edit it as your need.
 */
package httpsrv

import (
	"context"
	"encoding/json"
	"net/http"

	service "go-doudou-tutorials/microcomponent/component-b"

	"github.com/pkg/errors"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
)

type ComponentBHandlerImpl struct {
	componentB service.ComponentB
}

func (receiver *ComponentBHandlerImpl) Greeting(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx   context.Context
		msg   string
		reply string
		err   error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	if _, exists := _req.Form["msg"]; exists {
		msg = _req.FormValue("msg")
		if _err := rest.ValidateVar(msg, "", "msg"); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	} else {
		rest.HandleBadRequestErr(errors.New("missing parameter msg"))
	}
	reply, err = receiver.componentB.Greeting(
		ctx,
		msg,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Reply string `json:"reply"`
	}{
		Reply: reply,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}

func NewComponentBHandler(componentB service.ComponentB) ComponentBHandler {
	return &ComponentBHandlerImpl{
		componentB,
	}
}
