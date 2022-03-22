module statsvc

go 1.16

require (
	github.com/go-resty/resty/v2 v2.6.0
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
	github.com/unionj-cloud/go-doudou v1.0.2
)

require (
	github.com/apolloconfig/agollo/v4 v4.1.0
	github.com/gorilla/handlers v1.5.1 // indirect
)

replace github.com/unionj-cloud/go-doudou v1.0.2 => /Users/wubin1989/workspace/cloud/go-doudou
