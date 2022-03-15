package httpsrv

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	service "github.com/unionj-cloud/helloworld"
)

type HelloworldHandlerImpl struct {
	helloworld service.Helloworld
}

func (receiver *HelloworldHandlerImpl) Greeting(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx      context.Context
		greeting string
		data     string
		err      error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["greeting"]; exists {
		greeting = _req.FormValue("greeting")
	} else {
		http.Error(_writer, "missing parameter greeting", http.StatusBadRequest)
		return
	}
	data, err = receiver.helloworld.Greeting(
		ctx,
		greeting,
	)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
		} else if _err, ok := err.(*ddhttp.BizError); ok {
			http.Error(_writer, _err.Error(), _err.StatusCode)
		} else {
			http.Error(_writer, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data string `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewHelloworldHandler(helloworld service.Helloworld) HelloworldHandler {
	return &HelloworldHandlerImpl{
		helloworld,
	}
}

func (receiver *HelloworldHandlerImpl) Bye(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx      context.Context
		greeting string
		data     string
		err      error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["greeting"]; exists {
		greeting = _req.FormValue("greeting")
	} else {
		http.Error(_writer, "missing parameter greeting", http.StatusBadRequest)
		return
	}
	data, err = receiver.helloworld.Bye(
		ctx,
		greeting,
	)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
		} else if _err, ok := err.(*ddhttp.BizError); ok {
			http.Error(_writer, _err.Error(), _err.StatusCode)
		} else {
			http.Error(_writer, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data string `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusInternalServerError)
		return
	}
}
