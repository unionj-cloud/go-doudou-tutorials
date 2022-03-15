package httpsrv

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	service "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff/vo"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/toolkit/cast"
	v3 "github.com/unionj-cloud/go-doudou/toolkit/openapi/v3"
)

type WordcloudBffHandlerImpl struct {
	wordcloudBff service.WordcloudBff
}

func (receiver *WordcloudBffHandlerImpl) Upload(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		file v3.FileModel
		lang string
		top  *int
		data vo.UploadResult
		err  error
	)
	ctx = _req.Context()
	if _err := _req.ParseMultipartForm(32 << 20); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusBadRequest)
		return
	}
	fileFileHeaders, exists := _req.MultipartForm.File["file"]
	if exists {
		if len(fileFileHeaders) == 0 {
			http.Error(_writer, "no file uploaded for parameter file", http.StatusBadRequest)
			return
		}
		if len(fileFileHeaders) > 0 {
			_fh := fileFileHeaders[0]
			_f, _err := _fh.Open()
			if _err != nil {
				http.Error(_writer, _err.Error(), http.StatusBadRequest)
				return
			}
			file = v3.FileModel{
				Filename: _fh.Filename,
				Reader:   _f,
			}
		}
	} else {
		http.Error(_writer, "missing parameter file", http.StatusBadRequest)
		return
	}
	if _err := _req.ParseForm(); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["lang"]; exists {
		lang = _req.FormValue("lang")
	} else {
		http.Error(_writer, "missing parameter lang", http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["top"]; exists {
		if casted, err := cast.ToIntE(_req.FormValue("top")); err != nil {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
			return
		} else {
			top = &casted
		}
	}
	data, err = receiver.wordcloudBff.Upload(
		ctx,
		file,
		lang,
		top,
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
		Data vo.UploadResult `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewWordcloudBffHandler(wordcloudBff service.WordcloudBff) WordcloudBffHandler {
	return &WordcloudBffHandlerImpl{
		wordcloudBff,
	}
}

func (receiver *WordcloudBffHandlerImpl) GetTaskPage(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx      context.Context
		page     int
		pageSize int
		result   vo.TaskPageRet
		err      error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["page"]; exists {
		if casted, err := cast.ToIntE(_req.FormValue("page")); err != nil {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
			return
		} else {
			page = casted
		}
	} else {
		http.Error(_writer, "missing parameter page", http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["pageSize"]; exists {
		if casted, err := cast.ToIntE(_req.FormValue("pageSize")); err != nil {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
			return
		} else {
			pageSize = casted
		}
	} else {
		http.Error(_writer, "missing parameter pageSize", http.StatusBadRequest)
		return
	}
	result, err = receiver.wordcloudBff.GetTaskPage(
		ctx,
		page,
		pageSize,
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
		Result vo.TaskPageRet `json:"result"`
	}{
		Result: result,
	}); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusInternalServerError)
		return
	}
}
