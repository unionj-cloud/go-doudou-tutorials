package httpsrv

import (
	service "annotation"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
)

type AnnotationHandlerImpl struct {
	annotation service.Annotation
}

func (receiver *AnnotationHandlerImpl) GetGuest(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		data string
		err  error
	)
	ctx = _req.Context()
	data, err = receiver.annotation.GetGuest(
		ctx,
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
func (receiver *AnnotationHandlerImpl) GetUser(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		data string
		err  error
	)
	ctx = _req.Context()
	data, err = receiver.annotation.GetUser(
		ctx,
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
func (receiver *AnnotationHandlerImpl) GetAdmin(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		data string
		err  error
	)
	ctx = _req.Context()
	data, err = receiver.annotation.GetAdmin(
		ctx,
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

func NewAnnotationHandler(annotation service.Annotation) AnnotationHandler {
	return &AnnotationHandlerImpl{
		annotation,
	}
}
