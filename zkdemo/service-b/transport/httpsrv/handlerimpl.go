/**
* Generated by go-doudou v2.0.6.
* You can edit it as your need.
 */
package httpsrv

import (
	"context"
	"encoding/json"
	"net/http"
	service "service-b"
	"service-b/dto"

	"github.com/pkg/errors"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/cast"
)

type ServiceBHandlerImpl struct {
	serviceB service.ServiceB
}

func (receiver *ServiceBHandlerImpl) GetDeptById(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx    context.Context
		deptId int
		dept   dto.DeptDto
		err    error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		rest.HandleBadRequestErr(_err)
	}
	if _, exists := _req.Form["deptId"]; exists {
		if casted, _err := cast.ToIntE(_req.FormValue("deptId")); _err != nil {
			rest.HandleBadRequestErr(_err)
		} else {
			deptId = casted
		}
		if _err := rest.ValidateVar(deptId, "", "deptId"); _err != nil {
			rest.HandleBadRequestErr(_err)
		}
	} else {
		rest.HandleBadRequestErr(errors.New("missing parameter deptId"))
	}
	dept, err = receiver.serviceB.GetDeptById(
		ctx,
		deptId,
	)
	if err != nil {
		panic(err)
	}
	if _err := json.NewEncoder(_writer).Encode(struct {
		Dept dto.DeptDto `json:"dept"`
	}{
		Dept: dept,
	}); _err != nil {
		rest.HandleInternalServerError(_err)
	}
}

func NewServiceBHandler(serviceB service.ServiceB) ServiceBHandler {
	return &ServiceBHandlerImpl{
		serviceB,
	}
}
