FROM golang:1.16.6-alpine AS builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ARG user
ENV HOST_USER=$user

WORKDIR /repo

ADD ./wordcloud-user ./

RUN go mod tidy && go mod vendor

RUN export GDD_VER=$(go list -mod=vendor -m -f '{{ .Version }}' github.com/unionj-cloud/go-doudou) && \
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags="-X 'github.com/unionj-cloud/go-doudou/v2/framework/buildinfo.BuildUser=$HOST_USER' -X 'github.com/unionj-cloud/go-doudou/v2/framework/buildinfo.BuildTime=$(date)' -X 'github.com/unionj-cloud/go-doudou/v2/framework/buildinfo.GddVer=$GDD_VER'" -mod vendor -o api cmd/main.go

FROM alpine:3.14
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache bash tzdata curl

ENV TZ="Asia/Shanghai"

WORKDIR /repo

COPY --from=builder /repo/api ./

COPY /wordcloud-user/.env ./.env
COPY /wordcloud-user/.env.test ./.env.test

ENTRYPOINT ["/repo/api"]
