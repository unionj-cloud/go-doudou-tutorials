FROM devopsworks/golang-upx:1.18 AS builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ARG user
ENV HOST_USER=$user

WORKDIR /repo

ADD . ./

WORKDIR /repo/server

RUN go mod download
RUN go mod vendor

RUN export GDD_VER=$(go list -mod=vendor -m -f '{{ .Version }}' github.com/unionj-cloud/go-doudou) && \
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags="-s -w -X 'github.com/unionj-cloud/go-doudou/framework/buildinfo.BuildUser=$HOST_USER' -X 'github.com/unionj-cloud/go-doudou/framework/buildinfo.BuildTime=$(date)' -X 'github.com/unionj-cloud/go-doudou/framework/buildinfo.GddVer=$GDD_VER'" -mod vendor -o api cmd/main.go && \
strip api && /usr/local/bin/upx api

FROM alpine:3.14

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
# install curl command for docker-compose healthcheck
RUN apk add --no-cache curl

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ="Asia/Shanghai"

WORKDIR /repo

COPY --from=builder /repo/server/api ./server/

COPY server/.env* ./server/
COPY cert ./cert

WORKDIR /repo/server

ENTRYPOINT ["/repo/server/api"]
