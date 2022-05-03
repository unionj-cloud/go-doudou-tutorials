#!/bin/bash

export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct
export TZ="Asia/Shanghai"

go build -v -o api cmd/main.go

pm2 reload ecosystem.config.js --only usersvc --env $1
