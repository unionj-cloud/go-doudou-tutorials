package httpsrv

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	service "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-seg/vo"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
)

type WordcloudSegHandlerImpl struct {
	wordcloudSeg service.WordcloudSeg
}

func (receiver *WordcloudSegHandlerImpl) Seg(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx     context.Context
		payload vo.SegPayload
		rs      vo.SegResult
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
	rs, err = receiver.wordcloudSeg.Seg(
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
		Rs vo.SegResult `json:"rs"`
	}{
		Rs: rs,
	}); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewWordcloudSegHandler(wordcloudSeg service.WordcloudSeg) WordcloudSegHandler {
	return &WordcloudSegHandlerImpl{
		wordcloudSeg,
	}
}
