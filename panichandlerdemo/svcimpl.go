/**
* Generated by go-doudou v2.0.4.
* You can edit it as your need.
 */
package service

import (
	"context"
	"testsvc/config"

	"github.com/pkg/errors"

	"testsvc/dto"

	"github.com/brianvoe/gofakeit/v6"
)

var BookNotFoundException = errors.New("Book not found")
var ConversionFailedException = errors.New("Conversion failed")

var _ Testsvc = (*TestsvcImpl)(nil)

type TestsvcImpl struct {
	conf *config.Config
}

func NewTestsvc(conf *config.Config) *TestsvcImpl {
	return &TestsvcImpl{
		conf: conf,
	}
}

func (receiver *TestsvcImpl) GetBookNotFoundException(ctx context.Context) (re error) {
	panic(BookNotFoundException)
}
func (receiver *TestsvcImpl) GetConversionFailedException(ctx context.Context) (re error) {
	panic(ConversionFailedException)
}

func (receiver *TestsvcImpl) GetBookPage(ctx context.Context, name string, author string, page dto.Page) (re error) {
	var _result struct {
	}
	_ = gofakeit.Struct(&_result)
	return nil
}