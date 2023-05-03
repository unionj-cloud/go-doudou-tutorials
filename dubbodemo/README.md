# dubbodemo

## 工程结构

```shell
.
├── README.md
├── docker-compose.yml
├── dubbo
│   ├── go.mod
│   ├── go.sum
│   └── rpc
│       └── grpc
│           ├── go-client    # dubbo gRPC服务消费者
│           ├── go-server    # dubbo gRPC服务提供者
│           ├── protobuf
│           └── service-b
├── service-a                      # go-doudou RESTful服务a
└── service-b                      # go-doudou gRPC服务b
```

## 测试步骤
### 1. 启动zookeeper

`docker-compose -f docker-compose.yml up -d --remove-orphans`

### 2. 启动service-b

`cd service-b && go run cmd/main.go`

### 3. 启动go-server

`cd dubbo/rpc/grpc/go-server && go run cmd/server.go`

### 4. 启动service-a

`cd service-a && go run cmd/main.go`

### 5. 测试从go-doudou服务调用dubbo服务

```curl
curl --location 'http://localhost:6060/rpc/say/hello?name=wubin'
```

### 6. 测试从dubbo服务调用go-doudou服务

`cd dubbo/rpc/grpc/go-client && go run cmd/go-doudou/godoudou_client.go`