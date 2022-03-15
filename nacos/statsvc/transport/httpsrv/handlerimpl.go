package httpsrv

import (
	"context"
	"encoding/json"
	"net/http"
	service "statsvc"

	"github.com/pkg/errors"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/toolkit/cast"
)

type StatsvcHandlerImpl struct {
	statsvc service.Statsvc
}

func (receiver *StatsvcHandlerImpl) Add(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		x    int
		y    int
		data int
		err  error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["x"]; exists {
		if casted, err := cast.ToIntE(_req.FormValue("x")); err != nil {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
			return
		} else {
			x = casted
		}
	} else {
		http.Error(_writer, "missing parameter x", http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["y"]; exists {
		if casted, err := cast.ToIntE(_req.FormValue("y")); err != nil {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
			return
		} else {
			y = casted
		}
	} else {
		http.Error(_writer, "missing parameter y", http.StatusBadRequest)
		return
	}
	data, err = receiver.statsvc.Add(
		ctx,
		x,
		y,
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
		Data int `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewStatsvcHandler(statsvc service.Statsvc) StatsvcHandler {
	return &StatsvcHandlerImpl{
		statsvc,
	}
}

func (receiver *StatsvcHandlerImpl) GetEcho(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		s    string
		data string
		err  error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["s"]; exists {
		s = _req.FormValue("s")
	} else {
		http.Error(_writer, "missing parameter s", http.StatusBadRequest)
		return
	}
	data, err = receiver.statsvc.GetEcho(
		ctx,
		s,
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
