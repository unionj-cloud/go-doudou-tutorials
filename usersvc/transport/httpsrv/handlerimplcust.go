package httpsrv

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/toolkit/cast"
	"io"
	"net/http"
	"os"
)

func (receiver *UsersvcHandlerImpl) GetPublicDownloadAvatar(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx    context.Context
		userId int
		data   *os.File
		err    error
	)
	ctx = _req.Context()
	if err := _req.ParseForm(); err != nil {
		http.Error(_writer, err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["userId"]; exists {
		if casted, err := cast.ToIntE(_req.FormValue("userId")); err != nil {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
			return
		} else {
			userId = casted
		}
	} else {
		http.Error(_writer, "missing parameter userId", http.StatusBadRequest)
		return
	}
	data, err = receiver.usersvc.GetPublicDownloadAvatar(
		ctx,
		userId,
	)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			http.Error(_writer, err.Error(), http.StatusBadRequest)
		} else if err, ok := err.(*ddhttp.BizError); ok {
			http.Error(_writer, err.Error(), err.StatusCode)
		} else {
			http.Error(_writer, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	if data == nil {
		http.Error(_writer, "No file returned", http.StatusInternalServerError)
		return
	}
	defer data.Close()
	var _fi os.FileInfo
	_fi, _err := data.Stat()
	if _err != nil {
		http.Error(_writer, _err.Error(), http.StatusInternalServerError)
		return
	}
	_writer.Header().Set("Content-Disposition", "attachment; filename="+_fi.Name())
	_writer.Header().Set("Content-Type", "application/octet-stream")
	_writer.Header().Set("Content-Length", fmt.Sprintf("%d", _fi.Size()))
	io.Copy(_writer, data)
}
