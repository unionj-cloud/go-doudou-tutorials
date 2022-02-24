package service

import (
	"github.com/pkg/errors"
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
)

var (
	errAvatarNotFound = ddhttp.NewBizError(errors.New("avatar not found"))
)
