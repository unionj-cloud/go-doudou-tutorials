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
	return `First, we enabled compression in general by setting enabled as true. Then, we specifically enabled JSON data compression by adding application/json to the list of mime-types. Finally, notice that we set min-response-size to 1,024 bytes long. This is because if we compress short amounts of data, we may produce bigger data than the original.

Often, proxies such as NGINX or web servers such as the Apache HTTP Server deliver the JSON data to other services or front-ends. Configuring JSON data compression in these tools is beyond the scope of this tutorial.

A previous tutorial on gzip tells us that gzip has various compression levels. Our code examples use gzip with the default Java compression level. Spring Boot, proxies, or web servers may get different compression results for the same JSON data.

If we use JSON as the serialization protocol to store data, we'll need to compress and decompress the data ourselves.`, nil
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
