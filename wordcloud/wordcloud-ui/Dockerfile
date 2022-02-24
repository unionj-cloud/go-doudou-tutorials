FROM node:16.14.0-alpine3.14 AS vueBuilder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache bash tzdata automake autoconf g++ make libpng-dev libtool file build-base nasm musl zlib-dev
ENV TZ="Asia/Shanghai"

# Set working directory
WORKDIR /app
# Copy all files from current directory to working dir in image
COPY . .
# install node modules and build assets
RUN yarn install --registry=https://registry.npm.taobao.org && yarn build

# nginx state for serving content
FROM nginx:alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add --no-cache bash tzdata
ENV TZ="Asia/Shanghai"

# Set working directory to nginx asset directory
WORKDIR /var/www
# Remove default nginx static assets
RUN rm -rf ./*
RUN rm -rf /etc/nginx/conf.d/*.conf
# Copy static assets from builder stage
COPY --from=vueBuilder /app/dist spxx
COPY nginx.conf /etc/nginx/nginx.conf
COPY sites-enabled /etc/nginx/conf.d
# Containers run nginx with global directives and daemon off
ENTRYPOINT ["nginx", "-g", "daemon off;"]
