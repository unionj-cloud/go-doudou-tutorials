FROM golang:1.16.6-alpine AS builder

ENV GO111MODULE=on
ARG user
ENV HOST_USER=$user
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /repo

ADD go.mod .
ADD go.sum .

ADD . ./

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache bash tzdata

ENV TZ="Asia/Shanghai"

RUN go mod vendor

RUN export GDD_VER=$(go list -mod=vendor -m -f '{{ .Version }}' github.com/unionj-cloud/go-doudou) && \
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags="-X 'github.com/unionj-cloud/go-doudou/svc/config.BuildUser=$HOST_USER' -X 'github.com/unionj-cloud/go-doudou/svc/config.BuildTime=$(date)' -X 'github.com/unionj-cloud/go-doudou/svc/config.GddVer=$GDD_VER'" -mod vendor -o api cmd/main.go

ENTRYPOINT ["/repo/api"]
