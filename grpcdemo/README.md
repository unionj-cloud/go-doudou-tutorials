# grpcdemo

This is a demo project for showing how to use `go-doudou` to develop grpc service.

## About Server

Server is providing both http server listening on port 8080 for RESTful service and grpc server listening on port 50051
for grpc service in single process when started. There is no any real world business logic.

## About client

We won't directly take test on server, instead we will test client to see how client talks to server.  

There are two endpoints for us to test:

- `http://localhost:6060/greeting/1`: test RESTful service of server
- `http://localhost:6060/greeting`: test grpc service of server

## How to Run

There are all source code including generated code by go-doudou CLI in this repo. `Makefile` provides 
some convenient commands for you to run server and client services.

### Run on Host

If you want to run them on host:

- Run server: `make server-tls`
- Run client: `make client-tls`

Don't forget to run `go mod tidy` under each of `client` and `server` directories to download dependencies at first.

### Run on Docker

If you want to run them on docker (you should install docker-compose at first if not yet):

1. Build server image: `make build-server-tls`
2. Build client image: `make build-client-tls`
3. Run all services: `make up`
4. Shutdown all services: `make down`

## How to test

You can use any http client tool or load testing tool like postman or wrk. 
Recommend to use [hey](https://github.com/rakyll/hey).

### On Host (Macos)

- Load test RESTful service:

```shell
➜  go-doudou-tutorials git:(master) ✗ hey -m POST -c 50 -n 10000 http://localhost:6060/greeting/1\?greeting\=Jack

Summary:
  Total:	0.6909 secs
  Slowest:	0.0257 secs
  Fastest:	0.0003 secs
  Average:	0.0034 secs
  Requests/sec:	14474.4803
  
  Total data:	220000 bytes
  Size/request:	22 bytes

Response time histogram:
  0.000 [1]	|
  0.003 [5690]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.005 [3186]	|■■■■■■■■■■■■■■■■■■■■■■
  0.008 [547]	|■■■■
  0.010 [127]	|■
  0.013 [178]	|■
  0.016 [72]	|■
  0.018 [151]	|■
  0.021 [35]	|
  0.023 [8]	|
  0.026 [5]	|


Latency distribution:
  10% in 0.0012 secs
  25% in 0.0019 secs
  50% in 0.0026 secs
  75% in 0.0037 secs
  90% in 0.0057 secs
  95% in 0.0093 secs
  99% in 0.0170 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0003 secs, 0.0257 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0026 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0032 secs
  resp wait:	0.0032 secs, 0.0002 secs, 0.0256 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0057 secs

Status code distribution:
  [200]	10000 responses
```

- Load test grpc service:

```shell
➜  go-doudou-tutorials git:(master) ✗ hey -m POST -c 50 -n 10000 http://localhost:6060/greeting\?greeting\=Jack

Summary:
  Total:	0.6576 secs
  Slowest:	0.0354 secs
  Fastest:	0.0004 secs
  Average:	0.0032 secs
  Requests/sec:	15206.9509
  
  Total data:	220000 bytes
  Size/request:	22 bytes

Response time histogram:
  0.000 [1]	|
  0.004 [8579]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.007 [871]	|■■■■
  0.011 [98]	|
  0.014 [126]	|■
  0.018 [153]	|■
  0.021 [70]	|
  0.025 [51]	|
  0.028 [1]	|
  0.032 [0]	|
  0.035 [50]	|


Latency distribution:
  10% in 0.0012 secs
  25% in 0.0017 secs
  50% in 0.0024 secs
  75% in 0.0032 secs
  90% in 0.0046 secs
  95% in 0.0088 secs
  99% in 0.0216 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0004 secs, 0.0354 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0019 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0035 secs
  resp wait:	0.0030 secs, 0.0003 secs, 0.0354 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0051 secs

Status code distribution:
  [200]	10000 responses
```

About 5% performance improvement if using grpc service.

### On Docker (Alpine)

- Load test RESTful service: 

```shell
➜  go-doudou-tutorials git:(master) ✗ hey -m POST -c 50 -n 10000 http://localhost:6060/greeting/1\?greeting\=Jack

Summary:
  Total:	2.6018 secs
  Slowest:	0.0535 secs
  Fastest:	0.0022 secs
  Average:	0.0126 secs
  Requests/sec:	3843.5603
  
  Total data:	220000 bytes
  Size/request:	22 bytes

Response time histogram:
  0.002 [1]	|
  0.007 [2198]	|■■■■■■■■■■■■■■■■■■■■■■■
  0.012 [3806]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.018 [2057]	|■■■■■■■■■■■■■■■■■■■■■■
  0.023 [1069]	|■■■■■■■■■■■
  0.028 [454]	|■■■■■
  0.033 [220]	|■■
  0.038 [112]	|■
  0.043 [54]	|■
  0.048 [23]	|
  0.054 [6]	|


Latency distribution:
  10% in 0.0057 secs
  25% in 0.0077 secs
  50% in 0.0107 secs
  75% in 0.0159 secs
  90% in 0.0217 secs
  95% in 0.0264 secs
  99% in 0.0368 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0022 secs, 0.0535 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0018 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0022 secs
  resp wait:	0.0125 secs, 0.0021 secs, 0.0535 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0022 secs

Status code distribution:
  [200]	10000 responses
```

- Load test grpc service: 

```shell
➜  go-doudou-tutorials git:(master) ✗ hey -m POST -c 50 -n 10000 http://localhost:6060/greeting\?greeting\=Jack

Summary:
  Total:	2.2613 secs
  Slowest:	0.0397 secs
  Fastest:	0.0016 secs
  Average:	0.0111 secs
  Requests/sec:	4422.2262
  
  Total data:	220000 bytes
  Size/request:	22 bytes

Response time histogram:
  0.002 [1]	|
  0.005 [268]	|■■■
  0.009 [3391]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.013 [3865]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.017 [1639]	|■■■■■■■■■■■■■■■■■
  0.021 [541]	|■■■■■■
  0.025 [142]	|■
  0.028 [110]	|■
  0.032 [37]	|
  0.036 [3]	|
  0.040 [3]	|


Latency distribution:
  10% in 0.0068 secs
  25% in 0.0082 secs
  50% in 0.0104 secs
  75% in 0.0130 secs
  90% in 0.0163 secs
  95% in 0.0185 secs
  99% in 0.0265 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0000 secs, 0.0016 secs, 0.0397 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0014 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0021 secs
  resp wait:	0.0109 secs, 0.0016 secs, 0.0383 secs
  resp read:	0.0001 secs, 0.0000 secs, 0.0028 secs

Status code distribution:
  [200]	10000 responses
```

About 15% performance improvement if using grpc service.

**NOTE:** Though I compared performance of RESTful and grpc above, RESTful and grpc can't be compared with each other in terms of resource usage, 
they are totally different. RESTful uses multiple connections between client and server, while grpc uses one connection only.

## How to develop

If you want to play with this project deeply or use it as scaffold to develop real world project, it is very easy and convenient.
You just need to `cd server` or `cd client`, modify the interface in `svc.go` file by adding new method or change old method signatures, 
and run `go generate ./...`, that's all. `go-doudou` will make its best effort to generate code for you, you just need to put your 
own business logic code into stub methods which implementing the interface.  

The command `go generate ./...` will run two go-doudou commands under the hood: 

- `go-doudou svc http --handler --doc -c`: generate code for RESTful service
- `go-doudou svc grpc`: generate code for grpc service

Both of two are very safe, you can feel free to run them repeatedly once you made changes in `svc.go` file.

## Conclusion

`go-doudou` has been supported developing grpc service from v1.3.2 release. You can develop grpc server based on `go-doudou`.  

If you have already used `go-doudou` to develop a RESTful service and you want to add grpc server to it, you can install `go-doudou` v1.3.2, 
latest protobuf compiler `protoc`, latest `protoc-gen-go` and latest `protoc-gen-go-grpc`, and then run command `go-doudou svc grpc` 
under the root directory of the service, `go-doudou` will generate proto file, grpc server/client stubs, and some updates to `svcimpl.go` file for you, 
and finally you should add some code (will be explained in next post) to `main.go` file to run grpc server.  

If you want to develop a grpc service for a new project, you can just run 4 commands to start your service:  

1. `go-doudou svc init moduleName`
2. `go-doudou svc grpc`
3. `go mod tidy`
4. `go run cmd/main.go`

In step 2, `go-doudou` will generate `main.go` file for you.