module github.com/unionj-cloud/helloworld

go 1.16

require (
	github.com/brianvoe/gofakeit/v6 v6.19.0
	github.com/go-resty/resty/v2 v2.7.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/iancoleman/strcase v0.1.3
	github.com/jmoiron/sqlx v1.3.1
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/opentracing-contrib/go-stdlib v1.0.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.0
	github.com/sirupsen/logrus v1.8.1
	github.com/slok/goresilience v0.2.0
	github.com/unionj-cloud/go-doudou v1.2.11
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
)

replace github.com/unionj-cloud/go-doudou v1.2.11 => /Users/wubin1989/workspace/cloud/go-doudou
