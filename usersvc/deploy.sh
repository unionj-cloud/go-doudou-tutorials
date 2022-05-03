#!/bin/bash

# 设置编译程序所需环境变量
export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct
# 编译程序，生成可执行文件
go build -v -o api cmd/main.go

# 设置时区
export TZ="Asia/Shanghai"

# 通过pm2启动服务进程
pm2 restart ecosystem.config.js --only usersvc --env $1
