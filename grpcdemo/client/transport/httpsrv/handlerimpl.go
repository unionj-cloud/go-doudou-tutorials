package httpsrv

import (
	"context"
	"encoding/json"
	service "github.com/unionj-cloud/go-doudou-tutorials/grpcdemo/client"
	"github.com/unionj-cloud/go-doudou-tutorials/grpcdemo/client/vo"
	"net/http"

	"github.com/pkg/errors"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
)

type EnumDemoHandlerImpl struct {
	enumDemo service.EnumDemo
}

func (receiver *EnumDemoHandlerImpl) GetKeyboard(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx    context.Context
		layout vo.KeyboardLayout
		data   string
		err    error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["layout"]; exists {
		layout.StringSetter(_req.FormValue("layout"))
		if _err := ddhttp.ValidateVar(layout, "", "layout"); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(_writer, "missing parameter layout", http.StatusBadRequest)
		return
	}
	data, err = receiver.enumDemo.GetKeyboard(
		ctx,
		layout,
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
func (receiver *EnumDemoHandlerImpl) GetKeyboard2(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx    context.Context
		layout *vo.KeyboardLayout
		data   string
		err    error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["layout"]; exists {
		layout = new(vo.KeyboardLayout)
		layout.StringSetter(_req.FormValue("layout"))
		if _err := ddhttp.ValidateVar(layout, "", "layout"); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		}
	}
	data, err = receiver.enumDemo.GetKeyboard2(
		ctx,
		layout,
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
func (receiver *EnumDemoHandlerImpl) GetKeyboards(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx    context.Context
		layout []vo.KeyboardLayout
		data   []string
		err    error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["layout"]; exists {
		for _, item := range _req.Form["layout"] {
			var _layout vo.KeyboardLayout
			_layout.StringSetter(item)
			layout = append(layout, _layout)
		}
		if _err := ddhttp.ValidateVar(layout, "", "layout"); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		if _, exists := _req.Form["layout[]"]; exists {
			for _, item := range _req.Form["layout[]"] {
				var _layout vo.KeyboardLayout
				_layout.StringSetter(item)
				layout = append(layout, _layout)
			}
			if _err := ddhttp.ValidateVar(layout, "", "layout"); _err != nil {
				http.Error(_writer, _err.Error(), http.StatusBadRequest)
				return
			}
		} else {
			http.Error(_writer, "missing parameter layout", http.StatusBadRequest)
			return
		}
	}
	data, err = receiver.enumDemo.GetKeyboards(
		ctx,
		layout,
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
		Data []string `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusInternalServerError)
		return
	}
}
func (receiver *EnumDemoHandlerImpl) GetKeyboards2(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx    context.Context
		layout *[]vo.KeyboardLayout
		data   []string
		err    error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["layout"]; exists {
		layout = new([]vo.KeyboardLayout)
		for _, item := range _req.Form["layout"] {
			var _layout vo.KeyboardLayout
			_layout.StringSetter(item)
			*layout = append(*layout, _layout)
		}
		if _err := ddhttp.ValidateVar(layout, "", "layout"); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		if _, exists := _req.Form["layout[]"]; exists {
			layout = new([]vo.KeyboardLayout)
			for _, item := range _req.Form["layout[]"] {
				var _layout vo.KeyboardLayout
				_layout.StringSetter(item)
				*layout = append(*layout, _layout)
			}
			if _err := ddhttp.ValidateVar(layout, "", "layout"); _err != nil {
				http.Error(_writer, _err.Error(), http.StatusBadRequest)
				return
			}
		}
	}
	data, err = receiver.enumDemo.GetKeyboards2(
		ctx,
		layout,
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
		Data []string `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusInternalServerError)
		return
	}
}
func (receiver *EnumDemoHandlerImpl) GetKeyboards5(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx    context.Context
		layout = new([]vo.KeyboardLayout)
		data   []string
		err    error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["layout"]; exists {
		for _, item := range _req.Form["layout"] {
			var _layout vo.KeyboardLayout
			_layout.StringSetter(item)
			*layout = append(*layout, _layout)
		}
		if _err := ddhttp.ValidateVar(layout, "", "layout"); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		if _, exists := _req.Form["layout[]"]; exists {
			for _, item := range _req.Form["layout[]"] {
				var _layout vo.KeyboardLayout
				_layout.StringSetter(item)
				*layout = append(*layout, _layout)
			}
			if _err := ddhttp.ValidateVar(layout, "", "layout"); _err != nil {
				http.Error(_writer, _err.Error(), http.StatusBadRequest)
				return
			}
		}
	}
	data, err = receiver.enumDemo.GetKeyboards5(
		ctx,
		*layout...,
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
		Data []string `json:"data"`
	}{
		Data: data,
	}); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusInternalServerError)
		return
	}
}
func (receiver *EnumDemoHandlerImpl) Keyboard(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx      context.Context
		keyboard vo.Keyboard
		data     string
		err      error
	)
	ctx = _req.Context()
	if _req.Body == nil {
		http.Error(_writer, "missing request body", http.StatusBadRequest)
		return
	} else {
		if _err := json.NewDecoder(_req.Body).Decode(&keyboard); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		} else {
			if _err := ddhttp.ValidateStruct(keyboard); _err != nil {
				http.Error(_writer, _err.Error(), http.StatusBadRequest)
				return
			}
		}
	}
	data, err = receiver.enumDemo.Keyboard(
		ctx,
		keyboard,
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

func NewEnumDemoHandler(enumDemo service.EnumDemo) EnumDemoHandler {
	return &EnumDemoHandlerImpl{
		enumDemo,
	}
}

func (receiver *EnumDemoHandlerImpl) Greeting(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx      context.Context
		greeting string
		data     string
		err      error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["greeting"]; exists {
		greeting = _req.FormValue("greeting")
		if _err := ddhttp.ValidateVar(greeting, "", "greeting"); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(_writer, "missing parameter greeting", http.StatusBadRequest)
		return
	}
	data, err = receiver.enumDemo.Greeting(
		ctx,
		greeting,
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

func (receiver *EnumDemoHandlerImpl) Greeting1(_writer http.ResponseWriter, _req *http.Request) {
	var (
		ctx      context.Context
		greeting string
		data     string
		err      error
	)
	ctx = _req.Context()
	if _err := _req.ParseForm(); _err != nil {
		http.Error(_writer, _err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := _req.Form["greeting"]; exists {
		greeting = _req.FormValue("greeting")
		if _err := ddhttp.ValidateVar(greeting, "", "greeting"); _err != nil {
			http.Error(_writer, _err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(_writer, "missing parameter greeting", http.StatusBadRequest)
		return
	}
	data, err = receiver.enumDemo.Greeting1(
		ctx,
		greeting,
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
