FROM golang:1.18.8-alpine AS builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache build-base gcc abuild binutils binutils-doc gcc-doc bash tzdata

ENV TZ="Asia/Shanghai"

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ARG user
ENV HOST_USER=$user

WORKDIR /repo

ADD ./wordcloud-seg ./

RUN go mod tidy && \
    go build -v -ldflags="-X 'github.com/unionj-cloud/go-doudou/framework/buildinfo.BuildUser=$HOST_USER' -X 'github.com/unionj-cloud/go-doudou/framework/buildinfo.BuildTime=$(date)'" -o api cmd/main.go

# alpine
FROM alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache libstdc++ libstdc++6

WORKDIR /repo

COPY --from=builder /repo/api ./

ENV TZ="Asia/Shanghai"
ENV GDD_PORT=6063
ENV BIZ_MODEL_PATH="/repo/models"

COPY /wordcloud-seg/models.tar.gz ./models.tar.gz
COPY /wordcloud-seg/.env ./.env

RUN tar -xzf models.tar.gz && rm models.tar.gz

ENTRYPOINT ["/repo/api"]
