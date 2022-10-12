package httpsrv

import (
	"context"
	"encoding/json"
	"net/http"

	ddmodel "github.com/unionj-cloud/go-doudou/framework/http/model"

	"github.com/pkg/errors"
	service "github.com/unionj-cloud/go-doudou-tutorials/grpcdemo/server"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"

	"github.com/unionj-cloud/go-doudou-tutorials/grpcdemo/server/vo"
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
		} else if _err, ok := err.(*ddmodel.BizError); ok {
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
		} else if _err, ok := err.(*ddmodel.BizError); ok {
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

func (receiver *HelloworldHandlerImpl) BiStream(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx     context.Context
		stream  vo.Order
		stream1 vo.Page
		err     error
	)
	ctx = _req.Context()
	if _req.Body == nil {
		http.Error(_writer, "missing request body", http.StatusBadRequest)
		return
	} else {
		if _err := json.NewDecoder(_req.Body).Decode(&stream); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		} else {
			if _err := ddhttp.ValidateStruct(stream); _err != nil {
				http.Error(_writer, _err.Error(), http.StatusBadRequest)
				return
			}
		}
	}
	stream1, err = receiver.helloworld.BiStream(
		ctx,
		stream,
	)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
		} else if _err, ok := err.(*ddmodel.BizError); ok {
			http.Error(_writer, _err.Error(), _err.StatusCode)
		} else {
			http.Error(_writer, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Stream1 vo.Page `json:"stream1"`
	}{
		Stream1: stream1,
	}); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusInternalServerError)
		return
	}
}
func (receiver *HelloworldHandlerImpl) ClientStream(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx    context.Context
		stream vo.Order
		data   vo.Page
		err    error
	)
	ctx = _req.Context()
	if _req.Body == nil {
		http.Error(_writer, "missing request body", http.StatusBadRequest)
		return
	} else {
		if _err := json.NewDecoder(_req.Body).Decode(&stream); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		} else {
			if _err := ddhttp.ValidateStruct(stream); _err != nil {
				http.Error(_writer, _err.Error(), http.StatusBadRequest)
				return
			}
		}
	}
	data, err = receiver.helloworld.ClientStream(
		ctx,
		stream,
	)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
		} else if _err, ok := err.(*ddmodel.BizError); ok {
			http.Error(_writer, _err.Error(), _err.StatusCode)
		} else {
			http.Error(_writer, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data vo.Page `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusInternalServerError)
		return
	}
}
func (receiver *HelloworldHandlerImpl) ServerStream(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx     context.Context
		payload vo.Order
		stream  vo.Page
		err     error
	)
	ctx = _req.Context()
	if _req.Body == nil {
		http.Error(_writer, "missing request body", http.StatusBadRequest)
		return
	} else {
		if _err := json.NewDecoder(_req.Body).Decode(&payload); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		} else {
			if _err := ddhttp.ValidateStruct(payload); _err != nil {
				http.Error(_writer, _err.Error(), http.StatusBadRequest)
				return
			}
		}
	}
	stream, err = receiver.helloworld.ServerStream(
		ctx,
		payload,
	)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
		} else if _err, ok := err.(*ddmodel.BizError); ok {
			http.Error(_writer, _err.Error(), _err.StatusCode)
		} else {
			http.Error(_writer, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Stream vo.Page `json:"stream"`
	}{
		Stream: stream,
	}); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusInternalServerError)
		return
	}
}
