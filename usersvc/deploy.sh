#!/bin/bash

# 设置编译程序所需环境变量
# 开启go module
export GO111MODULE=on
# 设置goproxy，加快依赖下载速度
export GOPROXY=https://goproxy.cn,direct
# 编译程序，生成可执行文件
go build -v -o api cmd/main.go

# 因为本案例的代码里用到了go标准库time包，里面的time.Local静态变量会从TZ环境变量中取值，
# 如果没有配置此环境变量，则取默认值UTC时区，这通常不符合我们的需求
export TZ="Asia/Shanghai"

# 通过pm2启动服务进程，--only表示只部署usersvc应用，--env表示读取配置文件中ENV属性中的
# env_dev、env_prod、env_test中的哪一个，注意传参时不加env_前缀
# pm2 restart命令是在服务器上执行的命令，与pm2 deploy命令有本质区别
pm2 restart ecosystem.config.js --only usersvc --env $1