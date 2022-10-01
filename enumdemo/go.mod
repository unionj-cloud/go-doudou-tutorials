module abc2

go 1.13

require (
	github.com/go-resty/resty/v2 v2.7.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/iancoleman/strcase v0.2.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/klauspost/compress v1.15.9
	github.com/opentracing-contrib/go-stdlib v1.0.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.12.1 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/unionj-cloud/go-doudou v1.2.11
	github.com/unionj-cloud/helloworld v0.0.0
	google.golang.org/grpc v1.38.0
)

replace github.com/unionj-cloud/helloworld v0.0.0 => ../helloworld
