/**
* Generated by go-doudou v2.1.4.
* You can edit it as your need.
 */
package httpsrv

import (
	"context"
	"encoding/json"
	jsoncopier "github.com/unionj-cloud/go-doudou/v2/toolkit/copier"
	pb "go-doudou-tutorials/microcomponent/component-a/transport/grpc"
	"net/http"

	service "go-doudou-tutorials/microcomponent/component-a"
	"go-doudou-tutorials/microcomponent/component-a/dto"

	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest/httprouter"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/cast"
)

type ComponentAHttp2Grpc struct {
	componentA pb.ComponentAServiceServer
}

func (receiver *ComponentAHttp2Grpc) PostUser(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		user dto.GddUser
		data *pb.PostUserRpcResponse
		err  error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&user); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(user); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	var gddUser pb.GddUser
	jsoncopier.DeepCopy(user, &gddUser)
	data, err = receiver.componentA.PostUserRpc(
		ctx,
		&gddUser,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *ComponentAHttp2Grpc) GetUser_Id(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		id   int32
		data *pb.GddUser
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
	var rpcRequest pb.GetUserIdRpcRequest
	jsoncopier.DeepCopy(user, &rpcRequest)
	data, err = receiver.componentA.GetUserIdRpc(
		ctx,
		&rpcRequest,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(data); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *ComponentAHttp2Grpc) PutUser(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx  context.Context
		user dto.GddUser
		re   error
	)
	ctx = _req.Context()
	if _err := json.NewDecoder(_req.Body).Decode(&user); _err != nil {
		rest.HandleBadRequestErr(_err)
	} else {
		if _err := rest.ValidateStruct(user); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	}
	re = receiver.componentA.PutUser(
		ctx,
		user,
	)
	if re != nil {
		panic(re)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
	}{}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
func (receiver *ComponentAHttp2Grpc) DeleteUser_Id(_writer http.ResponseWriter, _req *http.Request) {
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
	re = receiver.componentA.DeleteUser_Id(
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
func (receiver *ComponentAHttp2Grpc) GetUsers(_writer http.ResponseWriter, _req *http.Request) {
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
	data, err = receiver.componentA.GetUsers(
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

func NewComponentAHandler(componentA service.ComponentA) ComponentAHandler {
	return &ComponentAHttp2Grpc{
		componentA,
	}
}

func (receiver *ComponentAHttp2Grpc) GetMyUsers(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx              context.Context
		parameterWrapper struct {
			Parameter dto.Parameter `form:"parameter"`
		}
		data dto.MyCustomPageResult
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
	data, err = receiver.componentA.GetMyUsers(
		ctx,
		parameterWrapper.Parameter,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Data dto.MyCustomPageResult `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}
