package httpsrv

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	service "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff/vo"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	v3 "github.com/unionj-cloud/go-doudou/toolkit/openapi/v3"
)

type WordcloudBffHandlerImpl struct {
	wordcloudBff service.WordcloudBff
}

func (receiver *WordcloudBffHandlerImpl) Upload(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		file v3.FileModel
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
	data, err = receiver.wordcloudBff.Upload(
		ctx,
		file,
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
