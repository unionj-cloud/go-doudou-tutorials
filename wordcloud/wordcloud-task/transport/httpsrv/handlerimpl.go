package httpsrv

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	service "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/vo"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
)

type WordcloudTaskHandlerImpl struct {
	wordcloudTask service.WordcloudTask
}

func (receiver *WordcloudTaskHandlerImpl) TaskSave(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx     context.Context
		payload vo.TaskPayload
		data    int
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
		}
	}
	data, err = receiver.wordcloudTask.TaskSave(
		ctx,
		payload,
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
func (receiver *WordcloudTaskHandlerImpl) TaskSuccess(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx     context.Context
		payload vo.TaskSuccess
		data    string
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
		}
	}
	data, err = receiver.wordcloudTask.TaskSuccess(
		ctx,
		payload,
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
func (receiver *WordcloudTaskHandlerImpl) TaskFail(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx     context.Context
		payload vo.TaskFail
		data    string
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
		}
	}
	data, err = receiver.wordcloudTask.TaskFail(
		ctx,
		payload,
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

func NewWordcloudTaskHandler(wordcloudTask service.WordcloudTask) WordcloudTaskHandler {
	return &WordcloudTaskHandlerImpl{
		wordcloudTask,
	}
}
