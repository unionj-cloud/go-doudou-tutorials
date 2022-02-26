FROM golang:1.16.6-alpine AS builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache build-base gcc abuild binutils binutils-doc gcc-doc bash tzdata

ENV TZ="Asia/Shanghai"

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ARG user
ENV HOST_USER=$user

WORKDIR /repo

ADD . ./

RUN go mod tidy && \
    go build -v -ldflags="-X 'github.com/unionj-cloud/go-doudou/framework/buildinfo.BuildUser=$HOST_USER' -X 'github.com/unionj-cloud/go-doudou/framework/buildinfo.BuildTime=$(date)' -X 'github.com/unionj-cloud/go-doudou/framework/buildinfo.GddVer=$GDD_VER'" -o api cmd/main.go

ENTRYPOINT ["/repo/api"]
