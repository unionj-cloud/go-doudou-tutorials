package service

import (
	"github.com/pkg/errors"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
)

var (
	errAvatarNotFound = rest.NewBizError(errors.New("avatar not found"))
)
