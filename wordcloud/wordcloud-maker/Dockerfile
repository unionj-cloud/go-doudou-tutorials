FROM golang:1.16.6 AS builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ARG user
ENV HOST_USER=$user

WORKDIR /repo

ADD . ./

RUN go mod tidy && \
go build -v -ldflags="-X 'github.com/unionj-cloud/go-doudou/framework/buildinfo.BuildUser=$HOST_USER' -X 'github.com/unionj-cloud/go-doudou/framework/buildinfo.BuildTime=$(date)' -X 'github.com/unionj-cloud/go-doudou/framework/buildinfo.GddVer=$GDD_VER'" -o api cmd/main.go

FROM ubuntu:16.04

RUN echo 'deb http://mirrors.aliyun.com/ubuntu/ xenial main \n\
deb-src http://mirrors.aliyun.com/ubuntu/ xenial main \n\
deb http://mirrors.aliyun.com/ubuntu/ xenial-updates main \n\
deb-src http://mirrors.aliyun.com/ubuntu/ xenial-updates main \n\
deb http://mirrors.aliyun.com/ubuntu/ xenial universe \n\
deb-src http://mirrors.aliyun.com/ubuntu/ xenial universe \n\
deb http://mirrors.aliyun.com/ubuntu/ xenial-updates universe \n\
deb-src http://mirrors.aliyun.com/ubuntu/ xenial-updates universe \n\
deb http://mirrors.aliyun.com/ubuntu/ xenial-security main \n\
deb-src http://mirrors.aliyun.com/ubuntu/ xenial-security main \n\
deb http://mirrors.aliyun.com/ubuntu/ xenial-security universe \n\
deb-src http://mirrors.aliyun.com/ubuntu/ xenial-security universe' > /etc/apt/sources.list

RUN apt-get update
RUN apt-get install -y language-pack-zh-hans curl
RUN echo 'LANG="zh_CN.UTF-8" \n\
LANGUAGE="zh_CN:zh:en_US:en"' >> /etc/environment
RUN echo 'en_US.UTF-8 UTF-8 \n\
zh_CN.UTF-8 UTF-8 \n\
zh_CN.GBK GBK \n\
zh_CN GB2312' >> /var/lib/locales/supported.d/local
RUN locale-gen
ENV LANG zh_CN.UTF-8
ENV LANGUAGE zh_CN:zh:en_US:en
ENV LC_ALL zh_CN.UTF-8
RUN apt-get install -y fonts-droid-fallback ttf-wqy-zenhei ttf-wqy-microhei fonts-arphic-ukai fonts-arphic-uming

# cn fonts
COPY fonts/simhei.ttf /usr/share/fonts/simhei.ttf

RUN curl -LO https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
RUN apt-get install -y ./google-chrome-stable_current_amd64.deb
RUN rm google-chrome-stable_current_amd64.deb

ENV TZ="Asia/Shanghai"

WORKDIR /repo

COPY --from=builder /repo/api ./

COPY .env ./.env

ENTRYPOINT ["/repo/api"]
