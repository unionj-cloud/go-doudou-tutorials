FROM devopsworks/golang-upx:1.18 AS builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ARG user
ENV HOST_USER=$user

WORKDIR /repo

# all the steps are cached
ADD go.mod .
ADD go.sum .
# if go.mod/go.sum not changed, this step is also cached
RUN go mod download

ADD . ./
RUN go mod vendor

RUN export GDD_VER=$(go list -mod=vendor -m -f '{{ .Version }}' github.com/unionj-cloud/go-doudou) && \
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags="-s -w -X 'github.com/unionj-cloud/go-doudou/framework/buildinfo.BuildUser=$HOST_USER' -X 'github.com/unionj-cloud/go-doudou/framework/buildinfo.BuildTime=$(date)' -X 'github.com/unionj-cloud/go-doudou/framework/buildinfo.GddVer=$GDD_VER'" -mod vendor -o api cmd/main.go && \
strip api && /usr/local/bin/upx api

FROM alpine:3.14

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
# install curl command for docker-compose healthcheck
RUN apk add --no-cache dumb-init

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ="Asia/Shanghai"

WORKDIR /repo

COPY --from=builder /repo/api ./

COPY .env* ./

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD /repo/api

#FROM ubuntu:16.04
#
#RUN apt-get update && apt-get install -y wget
#RUN wget https://github.com/Yelp/dumb-init/releases/download/v1.2.5/dumb-init_1.2.5_amd64.deb
#RUN dpkg -i dumb-init_*.deb
#
#ENV TZ="Asia/Shanghai"
#
#WORKDIR /repo
#
#COPY --from=builder /repo/api ./
#
#COPY .env* ./
#
#ENTRYPOINT ["dumb-init", "--"]
#CMD /repo/api