#!/bin/bash

# 设置编译程序所需环境变量
export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct
# 编译程序，生成可执行文件
go build -v -o api cmd/main.go

# 启动mysql实例
docker-compose -f docker-compose.yml up -d

# 等待mysql实例启动完毕，可以连接
sleep 15s

# 通过pm2启动服务进程
pm2 reload ecosystem.config.js --only usersvc --env $1
