package httpsrv

import (
	"context"
	"encoding/json"
	"net/http"
	service "usersvc"
	"usersvc/vo"

	"github.com/pkg/errors"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/toolkit/cast"
	v3 "github.com/unionj-cloud/go-doudou/toolkit/openapi/v3"
)

type UsersvcHandlerImpl struct {
	usersvc service.Usersvc
}

func (receiver *UsersvcHandlerImpl) PageUsers(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx   context.Context
		query vo.PageQuery
		data  vo.PageRet
		err   error
	)
	ctx = _req.Context()
	if _req.Body == nil {
		http.Error(_writer, "missing request body", http.StatusBadRequest)
		return
	} else {
		if _err := json.NewDecoder(_req.Body).Decode(&query); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		} else {
			if _err := ddhttp.ValidateStruct(query); _err != nil {
				http.Error(_writer, _err.Error(), http.StatusBadRequest)
				return
			}
		}
	}
	data, err = receiver.usersvc.PageUsers(
		ctx,
		query,
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
		Data vo.PageRet `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusInternalServerError)
		return
	}
}
func (receiver *UsersvcHandlerImpl) GetUser(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx    context.Context
		userId int
		data   vo.UserVo
		err    error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["userId"]; exists {
		if casted, _err := cast.ToIntE(_req.FormValue("userId")); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		} else {
			userId = casted
		}
		if _err := ddhttp.ValidateVar(userId, "", "userId"); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(_writer, "missing parameter userId", http.StatusBadRequest)
		return
	}
	data, err = receiver.usersvc.GetUser(
		ctx,
		userId,
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
		Data vo.UserVo `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusInternalServerError)
		return
	}
}
func (receiver *UsersvcHandlerImpl) PublicSignUp(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx      context.Context
		username string
		password string
		code     *string
		data     int
		err      error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["username"]; exists {
		username = _req.FormValue("username")
		if _err := ddhttp.ValidateVar(username, "", "username"); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(_writer, "missing parameter username", http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["password"]; exists {
		password = _req.FormValue("password")
		if _err := ddhttp.ValidateVar(password, "", "password"); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(_writer, "missing parameter password", http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["code"]; exists {
		_code := _req.FormValue("code")
		code = &_code
		if _err := ddhttp.ValidateVar(code, "", "code"); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		}
	}
	data, err = receiver.usersvc.PublicSignUp(
		ctx,
		username,
		password,
		code,
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
func (receiver *UsersvcHandlerImpl) PublicLogIn(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx      context.Context
		username string
		password string
		data     string
		err      error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["username"]; exists {
		username = _req.FormValue("username")
		if _err := ddhttp.ValidateVar(username, "", "username"); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(_writer, "missing parameter username", http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["password"]; exists {
		password = _req.FormValue("password")
		if _err := ddhttp.ValidateVar(password, "", "password"); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(_writer, "missing parameter password", http.StatusBadRequest)
		return
	}
	data, err = receiver.usersvc.PublicLogIn(
		ctx,
		username,
		password,
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
func (receiver *UsersvcHandlerImpl) UploadAvatar(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx    context.Context
		avatar v3.FileModel
		id     int
		data   string
		err    error
	)
	ctx = _req.Context()
	if _err := _req.ParseMultipartForm(32 << 20); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusBadRequest)
		return
	}
	avatarFileHeaders, exists := _req.MultipartForm.File["avatar"]
	if exists {
		if len(avatarFileHeaders) == 0 {
			http.Error(_writer, "no file uploaded for parameter avatar", http.StatusBadRequest)
			return
		}
		if len(avatarFileHeaders) > 0 {
			_fh := avatarFileHeaders[0]
			_f, _err := _fh.Open()
			if _err != nil {
				http.Error(_writer, _err.Error(), http.StatusBadRequest)
				return
			}
			avatar = v3.FileModel{
				Filename: _fh.Filename,
				Reader:   _f,
			}
		}
	} else {
		http.Error(_writer, "missing parameter avatar", http.StatusBadRequest)
		return
	}
	if _err := _req.ParseForm(); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["id"]; exists {
		if casted, _err := cast.ToIntE(_req.FormValue("id")); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		} else {
			id = casted
		}
		if _err := ddhttp.ValidateVar(id, "", "id"); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(_writer, "missing parameter id", http.StatusBadRequest)
		return
	}
	data, err = receiver.usersvc.UploadAvatar(
		ctx,
		avatar,
		id,
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

func NewUsersvcHandler(usersvc service.Usersvc) UsersvcHandler {
	return &UsersvcHandlerImpl{
		usersvc,
	}
}
