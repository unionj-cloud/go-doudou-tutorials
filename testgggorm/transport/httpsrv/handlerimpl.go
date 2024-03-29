/**
* Generated by go-doudou v2.1.1.
* You can edit it as your need.
 */
package httpsrv

import (
	"context"
	"encoding/json"
	"net/http"
	service "testgggorm"
	"testgggorm/dto"

	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest/httprouter"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/cast"
)

type TestgggormHandlerImpl struct {
	testgggorm service.Testgggorm
}

func (receiver *TestgggormHandlerImpl) PostTInvalidToken(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		body dto.TInvalidToken
		data int32
		err  error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&body); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(body); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	data, err = receiver.testgggorm.PostTInvalidToken(
		ctx,
		body,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data int32 `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) PostTInvalidTokens(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		body []dto.TInvalidToken
		data []int32
		err  error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&body); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateVar(body, "", ""); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	data, err = receiver.testgggorm.PostTInvalidTokens(
		ctx,
		body,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data []int32 `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) GetTInvalidToken_Id(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		id   int32
		data dto.TInvalidToken
		err  error
	)
	paramsFromCtx := httprouter.ParamsFromContext(_req.Context())
	ctx = _req.Context()
	if casted, _err := cast.ToInt32E(paramsFromCtx.ByName("id")); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		id = casted
	}
	if _err := rest.ValidateVar(id, "", "id"); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	data, err = receiver.testgggorm.GetTInvalidToken_Id(
		ctx,
		id,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data dto.TInvalidToken `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) PutTInvalidToken(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		body dto.TInvalidToken
		re   error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&body); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(body); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	re = receiver.testgggorm.PutTInvalidToken(
		ctx,
		body,
	)
	if re != nil {
		panic(re)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
	}{}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) DeleteTInvalidToken_Id(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx context.Context
		id  int32
		re  error
	)
	paramsFromCtx := httprouter.ParamsFromContext(_req.Context())
	ctx = _req.Context()
	if casted, _err := cast.ToInt32E(paramsFromCtx.ByName("id")); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		id = casted
	}
	if _err := rest.ValidateVar(id, "", "id"); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	re = receiver.testgggorm.DeleteTInvalidToken_Id(
		ctx,
		id,
	)
	if re != nil {
		panic(re)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
	}{}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) GetTInvalidTokens(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx              context.Context
		parameterWrapper struct {
			Parameter dto.Parameter `form:"parameter"`
		}
		data dto.Page
		err  error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	if _err := rest.DecodeForm(&parameterWrapper, _req.Form); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(parameterWrapper.Parameter); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	data, err = receiver.testgggorm.GetTInvalidTokens(
		ctx,
		parameterWrapper.Parameter,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data dto.Page `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) PostTJjj(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		body dto.TJjj
		data int32
		err  error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&body); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(body); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	data, err = receiver.testgggorm.PostTJjj(
		ctx,
		body,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data int32 `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) PostTJjjs(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		body []dto.TJjj
		data []int32
		err  error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&body); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateVar(body, "", ""); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	data, err = receiver.testgggorm.PostTJjjs(
		ctx,
		body,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data []int32 `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) GetTJjj_Id(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		id   int32
		data dto.TJjj
		err  error
	)
	paramsFromCtx := httprouter.ParamsFromContext(_req.Context())
	ctx = _req.Context()
	if casted, _err := cast.ToInt32E(paramsFromCtx.ByName("id")); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		id = casted
	}
	if _err := rest.ValidateVar(id, "", "id"); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	data, err = receiver.testgggorm.GetTJjj_Id(
		ctx,
		id,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data dto.TJjj `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) PutTJjj(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		body dto.TJjj
		re   error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&body); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(body); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	re = receiver.testgggorm.PutTJjj(
		ctx,
		body,
	)
	if re != nil {
		panic(re)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
	}{}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) DeleteTJjj_Id(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx context.Context
		id  int32
		re  error
	)
	paramsFromCtx := httprouter.ParamsFromContext(_req.Context())
	ctx = _req.Context()
	if casted, _err := cast.ToInt32E(paramsFromCtx.ByName("id")); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		id = casted
	}
	if _err := rest.ValidateVar(id, "", "id"); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	re = receiver.testgggorm.DeleteTJjj_Id(
		ctx,
		id,
	)
	if re != nil {
		panic(re)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
	}{}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) GetTJjjs(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx              context.Context
		parameterWrapper struct {
			Parameter dto.Parameter `form:"parameter"`
		}
		data dto.Page
		err  error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	if _err := rest.DecodeForm(&parameterWrapper, _req.Form); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(parameterWrapper.Parameter); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	data, err = receiver.testgggorm.GetTJjjs(
		ctx,
		parameterWrapper.Parameter,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data dto.Page `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) PostTPost(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		body dto.TPost
		data int32
		err  error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&body); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(body); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	data, err = receiver.testgggorm.PostTPost(
		ctx,
		body,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data int32 `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) PostTPosts(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		body []dto.TPost
		data []int32
		err  error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&body); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateVar(body, "", ""); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	data, err = receiver.testgggorm.PostTPosts(
		ctx,
		body,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data []int32 `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) GetTPost_Id(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		id   int32
		data dto.TPost
		err  error
	)
	paramsFromCtx := httprouter.ParamsFromContext(_req.Context())
	ctx = _req.Context()
	if casted, _err := cast.ToInt32E(paramsFromCtx.ByName("id")); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		id = casted
	}
	if _err := rest.ValidateVar(id, "", "id"); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	data, err = receiver.testgggorm.GetTPost_Id(
		ctx,
		id,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data dto.TPost `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) PutTPost(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		body dto.TPost
		re   error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&body); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(body); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	re = receiver.testgggorm.PutTPost(
		ctx,
		body,
	)
	if re != nil {
		panic(re)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
	}{}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) DeleteTPost_Id(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx context.Context
		id  int32
		re  error
	)
	paramsFromCtx := httprouter.ParamsFromContext(_req.Context())
	ctx = _req.Context()
	if casted, _err := cast.ToInt32E(paramsFromCtx.ByName("id")); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		id = casted
	}
	if _err := rest.ValidateVar(id, "", "id"); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	re = receiver.testgggorm.DeleteTPost_Id(
		ctx,
		id,
	)
	if re != nil {
		panic(re)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
	}{}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) GetTPosts(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx              context.Context
		parameterWrapper struct {
			Parameter dto.Parameter `form:"parameter"`
		}
		data dto.Page
		err  error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	if _err := rest.DecodeForm(&parameterWrapper, _req.Form); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(parameterWrapper.Parameter); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	data, err = receiver.testgggorm.GetTPosts(
		ctx,
		parameterWrapper.Parameter,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data dto.Page `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) PostTUser(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		body dto.TUser
		data int32
		err  error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&body); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(body); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	data, err = receiver.testgggorm.PostTUser(
		ctx,
		body,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data int32 `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) PostTUsers(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		body []dto.TUser
		data []int32
		err  error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&body); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateVar(body, "", ""); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	data, err = receiver.testgggorm.PostTUsers(
		ctx,
		body,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data []int32 `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) GetTUser_Id(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		id   int32
		data dto.TUser
		err  error
	)
	paramsFromCtx := httprouter.ParamsFromContext(_req.Context())
	ctx = _req.Context()
	if casted, _err := cast.ToInt32E(paramsFromCtx.ByName("id")); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		id = casted
	}
	if _err := rest.ValidateVar(id, "", "id"); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	data, err = receiver.testgggorm.GetTUser_Id(
		ctx,
		id,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data dto.TUser `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) PutTUser(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		body dto.TUser
		re   error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&body); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(body); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	re = receiver.testgggorm.PutTUser(
		ctx,
		body,
	)
	if re != nil {
		panic(re)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
	}{}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) DeleteTUser_Id(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx context.Context
		id  int32
		re  error
	)
	paramsFromCtx := httprouter.ParamsFromContext(_req.Context())
	ctx = _req.Context()
	if casted, _err := cast.ToInt32E(paramsFromCtx.ByName("id")); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		id = casted
	}
	if _err := rest.ValidateVar(id, "", "id"); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	re = receiver.testgggorm.DeleteTUser_Id(
		ctx,
		id,
	)
	if re != nil {
		panic(re)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
	}{}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) GetTUsers(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx              context.Context
		parameterWrapper struct {
			Parameter dto.Parameter `form:"parameter"`
		}
		data dto.Page
		err  error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	if _err := rest.DecodeForm(&parameterWrapper, _req.Form); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(parameterWrapper.Parameter); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	data, err = receiver.testgggorm.GetTUsers(
		ctx,
		parameterWrapper.Parameter,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data dto.Page `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) PostTWordCloudTask(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		body dto.TWordCloudTask
		data int32
		err  error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&body); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(body); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	data, err = receiver.testgggorm.PostTWordCloudTask(
		ctx,
		body,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data int32 `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) PostTWordCloudTasks(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		body []dto.TWordCloudTask
		data []int32
		err  error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&body); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateVar(body, "", ""); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	data, err = receiver.testgggorm.PostTWordCloudTasks(
		ctx,
		body,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data []int32 `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) GetTWordCloudTask_Id(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		id   int32
		data dto.TWordCloudTask
		err  error
	)
	paramsFromCtx := httprouter.ParamsFromContext(_req.Context())
	ctx = _req.Context()
	if casted, _err := cast.ToInt32E(paramsFromCtx.ByName("id")); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		id = casted
	}
	if _err := rest.ValidateVar(id, "", "id"); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	data, err = receiver.testgggorm.GetTWordCloudTask_Id(
		ctx,
		id,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data dto.TWordCloudTask `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) PutTWordCloudTask(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		body dto.TWordCloudTask
		re   error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&body); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(body); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	re = receiver.testgggorm.PutTWordCloudTask(
		ctx,
		body,
	)
	if re != nil {
		panic(re)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
	}{}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) DeleteTWordCloudTask_Id(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx context.Context
		id  int32
		re  error
	)
	paramsFromCtx := httprouter.ParamsFromContext(_req.Context())
	ctx = _req.Context()
	if casted, _err := cast.ToInt32E(paramsFromCtx.ByName("id")); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		id = casted
	}
	if _err := rest.ValidateVar(id, "", "id"); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	re = receiver.testgggorm.DeleteTWordCloudTask_Id(
		ctx,
		id,
	)
	if re != nil {
		panic(re)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
	}{}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) GetTWordCloudTasks(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx              context.Context
		parameterWrapper struct {
			Parameter dto.Parameter `form:"parameter"`
		}
		data dto.Page
		err  error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	if _err := rest.DecodeForm(&parameterWrapper, _req.Form); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(parameterWrapper.Parameter); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	data, err = receiver.testgggorm.GetTWordCloudTasks(
		ctx,
		parameterWrapper.Parameter,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data dto.Page `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) PostTAuthor(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		body dto.TAuthor
		data int32
		err  error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&body); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(body); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	data, err = receiver.testgggorm.PostTAuthor(
		ctx,
		body,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data int32 `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) PostTAuthors(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		body []dto.TAuthor
		data []int32
		err  error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&body); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateVar(body, "", ""); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	data, err = receiver.testgggorm.PostTAuthors(
		ctx,
		body,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data []int32 `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) GetTAuthor_Id(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		id   int32
		data dto.TAuthor
		err  error
	)
	paramsFromCtx := httprouter.ParamsFromContext(_req.Context())
	ctx = _req.Context()
	if casted, _err := cast.ToInt32E(paramsFromCtx.ByName("id")); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		id = casted
	}
	if _err := rest.ValidateVar(id, "", "id"); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	data, err = receiver.testgggorm.GetTAuthor_Id(
		ctx,
		id,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data dto.TAuthor `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) PutTAuthor(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		body dto.TAuthor
		re   error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&body); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(body); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	re = receiver.testgggorm.PutTAuthor(
		ctx,
		body,
	)
	if re != nil {
		panic(re)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
	}{}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) DeleteTAuthor_Id(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx context.Context
		id  int32
		re  error
	)
	paramsFromCtx := httprouter.ParamsFromContext(_req.Context())
	ctx = _req.Context()
	if casted, _err := cast.ToInt32E(paramsFromCtx.ByName("id")); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		id = casted
	}
	if _err := rest.ValidateVar(id, "", "id"); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	re = receiver.testgggorm.DeleteTAuthor_Id(
		ctx,
		id,
	)
	if re != nil {
		panic(re)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
	}{}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) GetTAuthors(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx              context.Context
		parameterWrapper struct {
			Parameter dto.Parameter `form:"parameter"`
		}
		data dto.Page
		err  error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	if _err := rest.DecodeForm(&parameterWrapper, _req.Form); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(parameterWrapper.Parameter); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	data, err = receiver.testgggorm.GetTAuthors(
		ctx,
		parameterWrapper.Parameter,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data dto.Page `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) PostTClient(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		body dto.TClient
		data int32
		err  error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&body); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(body); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	data, err = receiver.testgggorm.PostTClient(
		ctx,
		body,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data int32 `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) PostTClients(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		body []dto.TClient
		data []int32
		err  error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&body); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateVar(body, "", ""); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	data, err = receiver.testgggorm.PostTClients(
		ctx,
		body,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data []int32 `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) GetTClient_Id(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		id   int32
		data dto.TClient
		err  error
	)
	paramsFromCtx := httprouter.ParamsFromContext(_req.Context())
	ctx = _req.Context()
	if casted, _err := cast.ToInt32E(paramsFromCtx.ByName("id")); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		id = casted
	}
	if _err := rest.ValidateVar(id, "", "id"); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	data, err = receiver.testgggorm.GetTClient_Id(
		ctx,
		id,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data dto.TClient `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) PutTClient(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		body dto.TClient
		re   error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&body); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(body); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	re = receiver.testgggorm.PutTClient(
		ctx,
		body,
	)
	if re != nil {
		panic(re)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
	}{}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) DeleteTClient_Id(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx context.Context
		id  int32
		re  error
	)
	paramsFromCtx := httprouter.ParamsFromContext(_req.Context())
	ctx = _req.Context()
	if casted, _err := cast.ToInt32E(paramsFromCtx.ByName("id")); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		id = casted
	}
	if _err := rest.ValidateVar(id, "", "id"); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	re = receiver.testgggorm.DeleteTClient_Id(
		ctx,
		id,
	)
	if re != nil {
		panic(re)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
	}{}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) GetTClients(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx              context.Context
		parameterWrapper struct {
			Parameter dto.Parameter `form:"parameter"`
		}
		data dto.Page
		err  error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	if _err := rest.DecodeForm(&parameterWrapper, _req.Form); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(parameterWrapper.Parameter); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	data, err = receiver.testgggorm.GetTClients(
		ctx,
		parameterWrapper.Parameter,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data dto.Page `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) PostTDept(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		body dto.TDept
		data int32
		err  error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&body); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(body); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	data, err = receiver.testgggorm.PostTDept(
		ctx,
		body,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data int32 `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) PostTDepts(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		body []dto.TDept
		data []int32
		err  error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&body); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateVar(body, "", ""); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	data, err = receiver.testgggorm.PostTDepts(
		ctx,
		body,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data []int32 `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) GetTDept_Id(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		id   int32
		data dto.TDept
		err  error
	)
	paramsFromCtx := httprouter.ParamsFromContext(_req.Context())
	ctx = _req.Context()
	if casted, _err := cast.ToInt32E(paramsFromCtx.ByName("id")); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		id = casted
	}
	if _err := rest.ValidateVar(id, "", "id"); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	data, err = receiver.testgggorm.GetTDept_Id(
		ctx,
		id,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data dto.TDept `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) PutTDept(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		body dto.TDept
		re   error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&body); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(body); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	re = receiver.testgggorm.PutTDept(
		ctx,
		body,
	)
	if re != nil {
		panic(re)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
	}{}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) DeleteTDept_Id(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx context.Context
		id  int32
		re  error
	)
	paramsFromCtx := httprouter.ParamsFromContext(_req.Context())
	ctx = _req.Context()
	if casted, _err := cast.ToInt32E(paramsFromCtx.ByName("id")); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		id = casted
	}
	if _err := rest.ValidateVar(id, "", "id"); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	re = receiver.testgggorm.DeleteTDept_Id(
		ctx,
		id,
	)
	if re != nil {
		panic(re)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
	}{}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *TestgggormHandlerImpl) GetTDepts(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx              context.Context
		parameterWrapper struct {
			Parameter dto.Parameter `form:"parameter"`
		}
		data dto.Page
		err  error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	if _err := rest.DecodeForm(&parameterWrapper, _req.Form); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(parameterWrapper.Parameter); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	data, err = receiver.testgggorm.GetTDepts(
		ctx,
		parameterWrapper.Parameter,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data dto.Page `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}

func NewTestgggormHandler(testgggorm service.Testgggorm) TestgggormHandler {
	return &TestgggormHandlerImpl{
		testgggorm,
	}
}

func (receiver *TestgggormHandlerImpl) TAuthorPosts(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		body dto.SaveAuthorReqDTO
		err  error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&body); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(body); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	err = receiver.testgggorm.TAuthorPosts(
		ctx,
		body,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
	}{}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
