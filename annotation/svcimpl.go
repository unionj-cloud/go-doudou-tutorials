package service

import (
	"annotation/config"
	"context"

	"github.com/brianvoe/gofakeit/v6"
)

type AnnotationImpl struct {
	conf *config.Config
}

func (receiver *AnnotationImpl) GetGuest(ctx context.Context) (data string, err error) {
	return "ZpkSrYzH", nil
}
func (receiver *AnnotationImpl) GetUser(ctx context.Context) (data string, err error) {
	var _result struct {
		Data string
	}
	_ = gofakeit.Struct(&_result)
	return _result.Data, nil
}
func (receiver *AnnotationImpl) GetAdmin(ctx context.Context) (data string, err error) {
	var _result struct {
		Data string
	}
	_ = gofakeit.Struct(&_result)
	return _result.Data, nil
}

func NewAnnotation(conf *config.Config) Annotation {
	return &AnnotationImpl{
		conf,
	}
}
