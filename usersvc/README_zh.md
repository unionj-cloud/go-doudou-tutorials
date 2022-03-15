# 快速上手go-doudou开发单体RESTful服务

[EN](./README.md) [中文](./README_zh.md)

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
### TOC

- [需求清单](#%E9%9C%80%E6%B1%82%E6%B8%85%E5%8D%95)
- [学习目标](#%E5%AD%A6%E4%B9%A0%E7%9B%AE%E6%A0%87)
- [开发环境准备](#%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83%E5%87%86%E5%A4%87)
- [安装go-doudou](#%E5%AE%89%E8%A3%85go-doudou)
- [初始化工程](#%E5%88%9D%E5%A7%8B%E5%8C%96%E5%B7%A5%E7%A8%8B)
- [定义接口](#%E5%AE%9A%E4%B9%89%E6%8E%A5%E5%8F%A3)
- [生成代码](#%E7%94%9F%E6%88%90%E4%BB%A3%E7%A0%81)
- [启动服务](#%E5%90%AF%E5%8A%A8%E6%9C%8D%E5%8A%A1)
- [数据库和表结构准备](#%E6%95%B0%E6%8D%AE%E5%BA%93%E5%92%8C%E8%A1%A8%E7%BB%93%E6%9E%84%E5%87%86%E5%A4%87)
- [生成domain和dao层代码](#%E7%94%9F%E6%88%90domain%E5%92%8Cdao%E5%B1%82%E4%BB%A3%E7%A0%81)
- [用户注册接口](#%E7%94%A8%E6%88%B7%E6%B3%A8%E5%86%8C%E6%8E%A5%E5%8F%A3)
  - [修改domain](#%E4%BF%AE%E6%94%B9domain)
  - [PublicSignUp方法实现](#publicsignup%E6%96%B9%E6%B3%95%E5%AE%9E%E7%8E%B0)
  - [Postman测试](#postman%E6%B5%8B%E8%AF%95)
- [用户登录接口](#%E7%94%A8%E6%88%B7%E7%99%BB%E5%BD%95%E6%8E%A5%E5%8F%A3)
  - [PublicLogIn方法实现](#publiclogin%E6%96%B9%E6%B3%95%E5%AE%9E%E7%8E%B0)
  - [Postman测试](#postman%E6%B5%8B%E8%AF%95-1)
- [上传头像接口](#%E4%B8%8A%E4%BC%A0%E5%A4%B4%E5%83%8F%E6%8E%A5%E5%8F%A3)
  - [修改domain](#%E4%BF%AE%E6%94%B9domain-1)
  - [修改.env配置](#%E4%BF%AE%E6%94%B9env%E9%85%8D%E7%BD%AE)
  - [JWT校验中间件](#jwt%E6%A0%A1%E9%AA%8C%E4%B8%AD%E9%97%B4%E4%BB%B6)
  - [UploadAvatar方法实现](#uploadavatar%E6%96%B9%E6%B3%95%E5%AE%9E%E7%8E%B0)
- [下载头像接口](#%E4%B8%8B%E8%BD%BD%E5%A4%B4%E5%83%8F%E6%8E%A5%E5%8F%A3)
  - [GetPublicDownloadAvatar方法实现](#getpublicdownloadavatar%E6%96%B9%E6%B3%95%E5%AE%9E%E7%8E%B0)
- [用户详情接口](#%E7%94%A8%E6%88%B7%E8%AF%A6%E6%83%85%E6%8E%A5%E5%8F%A3)
  - [GetUser方法实现](#getuser%E6%96%B9%E6%B3%95%E5%AE%9E%E7%8E%B0)
  - [Postman测试](#postman%E6%B5%8B%E8%AF%95-2)
- [用户分页接口](#%E7%94%A8%E6%88%B7%E5%88%86%E9%A1%B5%E6%8E%A5%E5%8F%A3)
  - [导入测试数据](#%E5%AF%BC%E5%85%A5%E6%B5%8B%E8%AF%95%E6%95%B0%E6%8D%AE)
  - [PageUsers方法实现](#pageusers%E6%96%B9%E6%B3%95%E5%AE%9E%E7%8E%B0)
  - [Postman测试](#postman%E6%B5%8B%E8%AF%95-3)
- [服务部署](#%E6%9C%8D%E5%8A%A1%E9%83%A8%E7%BD%B2)
- [总结](#%E6%80%BB%E7%BB%93)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

笔者2015年开始接触go语言并采用go语言从事web项目开发至今，先后用过beego、gin、grpc等框架。这些框架非常优秀，通过学习它们的源码，也学到了很多。笔者之前在公司一直是单打独斗，一个人就把前后端的活包了，用现成的框架其实也蛮好。只是后来带了团队，接了不少项目，开始接触和学习敏捷开发、项目管理等方面的理论和实践，发现前后端不同成员之间沟通和联调也是需要很多成本的，特别是如果前端同事完全不懂后端，后端同事完全不懂前端的情况下会遇到不少头疼的事。于是萌生了用go语言开发一套低代码的、易于快速开发的、同时方便前后端同事沟通和联调的微服务框架，这就是[go-doudou微服务框架](https://github.com/unionj-cloud/go-doudou)。go-doudou框架主要基于gorilla的mux路由库做RESTful接口的快速生成，基于hashicorp公司开源的memberlist库做服务注册与发现和故障检测，同时支持开发单体应用和微服务应用。本教程将通过一个用户管理服务的案例来分几篇文章介绍如何用go-doudou开发单体RESTful接口。

### 需求清单
- 用户注册
- 用户登录
- 用户详情
- 用户分页
- 上传头像
- 下载头像

### 学习目标
- 用户详情、用户分页和上传头像需要采用jwt做权限校验
- 用户注册、用户登录和下载头像接口可以公开访问，无须鉴权
- 提供在线接口文档
- 提供go语言客户端SDK
- 提供mock接口实现
- 实现真实业务逻辑
- go-doudou内建的ddl表结构同步工具
- go-doudou内建的dao层代码生成和使用

### 开发环境准备
- docker环境: 推荐下载安装docker官方的desktop软件，[官方安装文档地址](https://docs.docker.com/engine/install/)
- IDE：推荐goland，当然vscode也可以

### 安装go-doudou
- 配置goproxy.cn代理，加速依赖下载
```
export GOPROXY=https://goproxy.cn,direct
```
- 如果你用的go版本是1.16以下版本：
```
GO111MODULE=on  go get -v github.com/unionj-cloud/go-doudou@v1.0.0-beta1
```
如果你用的go是1.16及以上版本：
```
go get -v github.com/unionj-cloud/go-doudou@v1.0.0-beta1
```
- goproxy.cn的同步会延迟一些，如果执行以上命令失败，可以关闭代理，科学上网
```
export GOPROXY=https://proxy.golang.org,direct
```
- 以上办法都不行，可以直接克隆同步到gitee的源码，本地安装
```
git clone git@gitee.com:unionj-cloud/go-doudou.git
```
切到根路径下，执行命令：
```
go install
```
- 执行命令`go-doudou -v`，如果输出如下内容，表示安装成功：
```
➜  ~ go-doudou -v
go-doudou version v1.0.0-beta
```

### 初始化工程
执行命令：
```
go-doudou svc init usersvc 
```
切到usersvc路径下，可以看到生成了如下文件结构：
```
➜  tutorials ll
total 0
drwxr-xr-x  9 wubin1989  staff   288B 10 24 20:05 usersvc
➜  tutorials cd usersvc  
➜  usersvc git:(master) ✗ ll
total 24
-rw-r--r--  1 wubin1989  staff   707B 10 24 20:05 Dockerfile
-rw-r--r--  1 wubin1989  staff   439B 10 24 20:05 go.mod
-rw-r--r--  1 wubin1989  staff   247B 10 24 20:05 svc.go
drwxr-xr-x  3 wubin1989  staff    96B 10 24 20:05 vo
```
- svc.go文件：做接口设计和定义
- vo文件夹：定义接口入参和出参的结构体
- Dockerfile：用于docker镜像打包

### 定义接口
我们打开svc.go文件看一下：
```go
package service

import (
	"context"
	v3 "github.com/unionj-cloud/go-doudou/toolkit/openapi/v3"
	"os"
	"usersvc/vo"
)

// Usersvc 用户管理服务
// 调用用户详情、用户分页和上传头像接口需要带上Bearer Token请求头
// 用户注册、用户登录和下载头像接口可以公开访问，无须鉴权
// Usersvc is user management service
// You should set Bearer Token header when you request protected endpoints such as user detail, user pagination and upload avatar.
// You can add doc for whole service here
type Usersvc interface {
	// PageUsers 用户分页接口
	// 展示如何定义POST请求且Content-Type为application/json的接口
	// PageUsers is user pagination api
	// demo how to define post request api which accepts application/json content-type
	PageUsers(ctx context.Context,
		// 分页请求参数
		// pagination parameter
		query vo.PageQuery) (
		// 分页结果
		// pagination result
		data vo.PageRet,
		// 错误信息
		// error
		err error)

	// GetUser 用户详情接口
	// 展示如何定义带查询字符串参数的GET请求接口
	// GetUser is user detail api
	// demo how to define get http request with query string parameters
	GetUser(ctx context.Context,
		// 用户ID
		// user id
		userId int) (
		// 用户详情
		// user detail
		data vo.UserVo,
		// 错误信息
		// error
		err error)

	// PublicSignUp 用户注册接口
	// 展示如何定义POST请求且Content-Type是application/x-www-form-urlencoded的接口
	// PublicSignUp is user signup api
	// demo how to define post request api which accepts application/x-www-form-urlencoded content-type
	PublicSignUp(ctx context.Context,
		// 用户名
		// username
		username string,
		// 密码
		// password
		password string,
		// 图形验证码
		// image code
		code *string,
	) (
		// 成功返回用户ID
		// return user ID if success
		data int, err error)

	// PublicLogIn 用户登录接口
	// 展示如何鉴权并返回token
	// PublicLogIn is user login api
	// demo how to do authentication and issue token
	PublicLogIn(ctx context.Context,
		// 用户名
		// username
		username string,
		// 密码
		// password
		password string) (
		// token
		data string, err error)

	// UploadAvatar 上传头像接口
	// 展示如何定义文件上传接口
	// 函数签名的入参里必须要有至少一个[]*v3.FileModel或者*v3.FileModel类型的参数
	// UploadAvatar is avatar upload api
	// demo how to define file upload api
	// NOTE: there must be at least one []*v3.FileModel or *v3.FileModel input parameter
	UploadAvatar(ctx context.Context,
		// 用户头像
		// user avatar
		avatar v3.FileModel, id int) (
		// 成功返回OK
		// return OK if success
		data string, err error)

	// GetPublicDownloadAvatar 下载头像接口
	// 展示如何定义文件下载接口
	// 函数签名的出参里必须有且只有一个*os.File类型的参数
	// GetPublicDownloadAvatar is avatar download api
	// demo how to define file download api
	// NOTE: there must be one and at most one *os.File output parameter
	GetPublicDownloadAvatar(ctx context.Context,
		// 用户ID
		// user id
		userId int) (
		// 文件二进制流
		// avatar file
		data *os.File, err error)
}
```
以上代码里每个方法都有注释。请仔细阅读。接口定义支持文档注释，只支持go语言常见的`//`注释。这些注释会作为OpenAPI3.0规范里的`description`参数值导出到生成的json文档和go-doudou内建的在线文档里，下文会做演示。

### 生成代码
执行如下命令，即可生成启动一个服务所需的全部代码
```
go-doudou svc http --handler -c go --doc
```
解释一下命令中的flag参数：
- --handler：表示需要生成http handler接口实现，就是把解析http请求参数和编码返回值的代码都生成出来
- -c：表示生成服务接口的客户端SDK，目前只支持`go`。如果不需要生成客户端SDK，可以不设置这个flag，因为相对其他代码来说，生成过程比较耗时
- --doc：表示生成OpenAPI3.0规范的json文档
  这行命令是笔者常用的命令，推荐大家也这样使用。并且这行命令可以在每次修改了svc.go文件里的接口定义以后执行，可以增量的生成代码。规则是：
- handler.go文件和OpenAPI3.0规范的json文档总是会重新生成
- handlerimpl.go文件和svcimpl.go文件只会增量生成，不会修改现有代码
- 其他文件都会先判断同名文件是否存在，如果存在就跳过

为了确保依赖都已经下载下来，最好再执行一下这个命令：
```
go mod tidy
```
我们来看一下此时的项目结构：
```
➜  usersvc git:(master) ✗ ll
total 296
-rw-r--r--  1 wubin1989  staff   707B 10 24 20:05 Dockerfile
drwxr-xr-x  3 wubin1989  staff    96B 10 24 23:10 client
drwxr-xr-x  3 wubin1989  staff    96B 10 24 23:10 cmd
drwxr-xr-x  3 wubin1989  staff    96B 10 24 23:10 config
drwxr-xr-x  3 wubin1989  staff    96B 10 24 23:10 db
-rw-r--r--  1 wubin1989  staff   514B 10 24 23:10 go.mod
-rw-r--r--  1 wubin1989  staff   115K 10 24 23:10 go.sum
-rw-r--r--  1 wubin1989  staff   1.7K 10 24 23:21 svc.go
-rw-r--r--  1 wubin1989  staff   1.6K 10 25 09:18 svcimpl.go
drwxr-xr-x  3 wubin1989  staff    96B 10 24 23:10 transport
-rwxr-xr-x  1 wubin1989  staff   5.9K 10 25 09:18 usersvc_openapi3.go
-rwxr-xr-x  1 wubin1989  staff   5.7K 10 25 09:18 usersvc_openapi3.json
drwxr-xr-x  3 wubin1989  staff    96B 10 24 23:07 vo
```
- Dockerfile文件：用于打包docker镜像
- client包：生成的go客户端代码
- cmd包：里面有main.go文件，用于启动服务
- config包：用于加载配置
- db包：用于连接数据库
- svc.go文件：设计接口
- svcimpl.go文件：里面有mock的接口实现，后续在里面根据业务需求编写真实的业务逻辑
- transport包：里面是http handler接口和实现，负责具体的接口入参解析和出参序列化
- usersvc_openapi3.go文件：用于在线接口文档功能
- usersvc_openapi3.json文件：遵循OpenAPI 3.0规范的接口文档
- vo包：里面是接口的入参和出参结构体类型

### 启动服务
```
go-doudou svc run
```
我们可以看到如下输出：
```
➜  usersvc git:(master) ✗ go-doudou svc run                       
INFO[2021-12-28 22:39:35] Initializing logging reporter                
INFO[2021-12-28 22:39:35] ================ Registered Routes ================ 
INFO[2021-12-28 22:39:35] +----------------------+--------+-------------------------+ 
INFO[2021-12-28 22:39:35] |         NAME         | METHOD |         PATTERN         | 
INFO[2021-12-28 22:39:35] +----------------------+--------+-------------------------+ 
INFO[2021-12-28 22:39:35] | PageUsers            | POST   | /page/users             | 
INFO[2021-12-28 22:39:35] | User                 | GET    | /user                   | 
INFO[2021-12-28 22:39:35] | PublicSignUp         | POST   | /public/sign/up         | 
INFO[2021-12-28 22:39:35] | PublicLogIn          | POST   | /public/log/in          | 
INFO[2021-12-28 22:39:35] | UploadAvatar         | POST   | /upload/avatar          | 
INFO[2021-12-28 22:39:35] | PublicDownloadAvatar | GET    | /public/download/avatar | 
INFO[2021-12-28 22:39:35] | GetDoc               | GET    | /go-doudou/doc          | 
INFO[2021-12-28 22:39:35] | GetOpenAPI           | GET    | /go-doudou/openapi.json | 
INFO[2021-12-28 22:39:35] | Prometheus           | GET    | /go-doudou/prometheus   | 
INFO[2021-12-28 22:39:35] | GetRegistry          | GET    | /go-doudou/registry     | 
INFO[2021-12-28 22:39:35] +----------------------+--------+-------------------------+ 
INFO[2021-12-28 22:39:35] =================================================== 
INFO[2021-12-28 22:39:35] Started in 233.424µs                         
INFO[2021-12-28 22:39:35] Http server is listening on :6060      
```
当出现"Http server is listening on :6060"时，表示服务已经启动，并且我们已经有了mock的服务接口实现。例如，我们可以执行如下命令请求`/user`接口，看看返回什么数据：
```
➜  usersvc git:(master) ✗ http http://localhost:6060/user
HTTP/1.1 200 OK
Content-Encoding: gzip
Content-Length: 109
Content-Type: application/json; charset=UTF-8
Date: Mon, 01 Nov 2021 15:21:10 GMT
Vary: Accept-Encoding

{
    "data": {
        "Dept": "ZkkCmcLU",
        "Id": -1941954111002502016,
        "Name": "aiMtQ",
        "Phone": "XMAqXf"
    }
}
```
此时你可能注意到返回的数据的字段名称是首字母大写的，这可能不是你想要的。在vo包下的vo.go文件里有一行`go generate`命令：
```
//go:generate go-doudou name --file $GOFILE
```
这行命令里用到了go-doudou框架内置的一个工具name。它可以根据指定的命名规则生成结构体字段后面的json标签。默认生成策略是首字母小写的驼峰命名策略，同时支持蛇形命名。未导出的字段会跳过，只修改导出字段的json标签。命令行执行命令：
```
go generate ./...
```
然后重启服务，请求`/user`接口，可以看到字段名称已经变成首字母小写的驼峰命名了。
```
➜  usersvc git:(master) ✗ http http://localhost:6060/user
HTTP/1.1 200 OK
Content-Encoding: gzip
Content-Length: 114
Content-Type: application/json; charset=UTF-8
Date: Tue, 02 Nov 2021 08:25:39 GMT
Vary: Accept-Encoding

{
    "data": {
        "dept": "wGAEEeveHp",
        "id": -816946940349962228,
        "name": "hquwOKl",
        "phone": "AriWmKYB"
    }
}
```
关于name工具的更多用法，请参考[文档](https://github.com/unionj-cloud/go-doudou/blob/main/name/README_zh.md)。
此时，因为vo包里的结构体修改了json标签，所以OpenAPI文档需要重新生成，否则在线文档里的字段名称还是修改前的。需要执行如下命令：
```
go-doudou svc http --doc
```
然后我们重启一下服务，在地址栏输入http://localhost:6060/go-doudou/doc， 再输入http basic用户名admin，密码admin，看一下在线文档是什么效果：

![image.png](https://p9-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/7ee9e355b00c49b1b7c95eb8125928d3~tplv-k3u1fbpfcp-watermark.image?)

![image.png](https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/a80f23feb5384ab9a0a17e1828917c0f~tplv-k3u1fbpfcp-watermark.image?)
在线文档里的接口说明和参数说明都取自svc.go的接口方法注释和参数注释。
<br/>
### 数据库和表结构准备
为了支持中文字符，需先在根目录下创建mysql配置文件my/custom.cnf，贴进去如下内容：
```
[client]
default-character-set=utf8mb4

[mysql]
default-character-set=utf8mb4

[mysqld]
character_set_server=utf8mb4
collation-server=utf8mb4_general_ci
default-authentication-plugin=mysql_native_password
init_connect='SET NAMES utf8mb4'
```
在根目录下创建数据库初始化脚本sqlscripts/init.sql，贴进去如下内容：
```sql
CREATE SCHEMA `tutorial` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

CREATE TABLE `tutorial`.`t_user`
(
    `id`        INT         NOT NULL AUTO_INCREMENT,
    `username`  VARCHAR(45) NOT NULL COMMENT '用户名',
    `password`  VARCHAR(60) NOT NULL COMMENT '密码',
    `name`      VARCHAR(45) NOT NULL COMMENT '真实姓名',
    `phone`     VARCHAR(45) NOT NULL COMMENT '手机号',
    `dept`      VARCHAR(45) NOT NULL COMMENT '所属部门',
    `create_at` DATETIME    NULL DEFAULT current_timestamp,
    `update_at` DATETIME    NULL DEFAULT current_timestamp on update current_timestamp,
    `delete_at` DATETIME    NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO tutorial.t_user (id, username, password, name, phone, dept, create_at, update_at, delete_at, avatar) VALUES (2, 'peter', '$2a$14$VaQLa/GbLAhRZvvTlgE8OOQgsBY4RDAJC5jkz13kjP9RlntdKBZVW', '张三丰', '13552053960', '技术部', '2021-12-28 06:41:00', '2021-12-28 14:59:20', null, 'out/wolf-wolves-snow-wolf-landscape-985ca149f06cd03b9f0ed8dfe326afdb.jpg');
INSERT INTO tutorial.t_user (id, username, password, name, phone, dept, create_at, update_at, delete_at, avatar) VALUES (4, 'john', '$2a$14$AKCs.u9vFUOCe5VwcmdfwOAkeiDtQYEgIB/nSU8/eemYwd91.qU.i', '李世民', '13552053961', '行政部', '2021-12-28 12:12:32', '2021-12-28 14:59:20', null, '');
INSERT INTO tutorial.t_user (id, username, password, name, phone, dept, create_at, update_at, delete_at, avatar) VALUES (5, 'lucy', '$2a$14$n0.l54axUqnKGagylQLu7ee.yDrtLubxzM1qmOaHK9Ft2P09YtQUS', '朱元璋', '13552053962', '销售部', '2021-12-28 12:13:17', '2021-12-28 14:59:20', null, '');
INSERT INTO tutorial.t_user (id, username, password, name, phone, dept, create_at, update_at, delete_at, avatar) VALUES (6, 'jack', '$2a$14$jFCwiZHcD7.DL/teao.Dl.HAFwk8wM2f1riH1fG2f52WYKqSiGZlC', '张无忌', '', '总裁办', '2021-12-28 12:14:19', '2021-12-28 14:59:20', null, '');
```
在根目录下创建docker-compose.yml文件，贴进入如下内容：
```
version: '3.9'

services:
  db:
    container_name: db
    image: mysql:5.7
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: 1234
    ports:
      - 3306:3306
    volumes:
      - $PWD/my:/etc/mysql/conf.d
      - $PWD/sqlscripts:/docker-entrypoint-initdb.d
    networks:
      - tutorial

networks:
  tutorial:
    driver: bridge
```
在根目录下执行docker compose命令，即可启动mysql数据库容器:
```
docker-compose -f docker-compose.yml up -d
```
可以通过`docker ps`命令查看正在运行的容器
```
➜  usersvc git:(master) ✗ docker ps        
CONTAINER ID   IMAGE       COMMAND                  CREATED          STATUS          PORTS                                                  NAMES
df6af6362c41   mysql:5.7   "docker-entrypoint.s…"   13 minutes ago   Up 13 minutes   0.0.0.0:3306->3306/tcp, :::3306->3306/tcp, 33060/tcp   db
```

### 生成domain和dao层代码
因为我们初始化的schema名称是tutorial，所以我们先要把.env文件里的环境变量DB_SCHEMA的值改成tutorial
```
DB_SCHEMA=tutorial
```
执行如下命令，生成domain和dao层代码：
```
go-doudou ddl -r --dao --pre=t_
```
解释一下：
- -r：表示从数据库表结构生成go结构体
- --dao：表示生成dao层代码
- --pre：表示表名称有前缀t_
  此时，你可以看到项目里多了两个目录：

![image.png](https://p9-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/9c364a17cf384f13bad7a79cf4a97852~tplv-k3u1fbpfcp-watermark.image?)
具体用法请参考[ddl文档](https://github.com/unionj-cloud/go-doudou/blob/main/ddl/doc/README.md)
这里我们看一下dao/base.go文件里提供了哪些CRUD方法，后面实现具体业务逻辑的时候会用到：
```go
package dao

import (
	"context"
	"github.com/unionj-cloud/go-doudou/toolkit/sqlext/query"
)

type Base interface {
	Insert(ctx context.Context, data interface{}) (int64, error)
	Upsert(ctx context.Context, data interface{}) (int64, error)
	UpsertNoneZero(ctx context.Context, data interface{}) (int64, error)
	DeleteMany(ctx context.Context, where query.Q) (int64, error)
	Update(ctx context.Context, data interface{}) (int64, error)
	UpdateNoneZero(ctx context.Context, data interface{}) (int64, error)
	UpdateMany(ctx context.Context, data interface{}, where query.Q) (int64, error)
	UpdateManyNoneZero(ctx context.Context, data interface{}, where query.Q) (int64, error)
	Get(ctx context.Context, id interface{}) (interface{}, error)
	SelectMany(ctx context.Context, where ...query.Q) (interface{}, error)
	CountMany(ctx context.Context, where ...query.Q) (int, error)
	PageMany(ctx context.Context, page query.Page, where ...query.Q) (query.PageRet, error)
}
```
再修改一下svcimpl.go文件的UsersvcImpl结构体
```go
type UsersvcImpl struct {
	conf *config.Config
	db   *sqlx.DB
}
```
以及NewUsersvc方法
```go
func NewUsersvc(conf *config.Config, db *sqlx.DB) Usersvc {
	return &UsersvcImpl{
		conf,
		db,
	}
}
```
生成的main方法里已经为我们注入了mysql连接实例，所以不用改
```go
svc := service.NewUsersvc(conf, conn)
```
后面我们直接在接口实现里面调用`UsersvcImpl`结构体的`db`属性即可

### 用户注册接口
#### 修改domain
因为通常来说用户名都必须是唯一的，所以我们需要改一下domain/user.go文件：
```go
Username string     `dd:"type:varchar(45);extra:comment '用户名';unique"`
```
再执行ddl命令
```shell
go-doudou ddl --pre=t_
```
这行命令没有-r参数了，表示从go结构体更新到表结构。

![image.png](https://p9-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/2412f44217bf4fd2a50d48b55521e90b~tplv-k3u1fbpfcp-watermark.image?)

#### PublicSignUp方法实现
要实现注册逻辑，我们需要先给dao层代码加一个方法CheckUsernameExists，判断一下传进来的用户名是否已经被注册。先改一下dao/userdao.go文件
```go
package dao

import "context"

type UserDao interface {
	Base
	CheckUsernameExists(ctx context.Context, username string) (bool, error)
}
```
再新建一个文件dao/userdaoimplext.go文件，加入如下代码
```go
package dao

import (
	"context"
	"github.com/unionj-cloud/go-doudou/ddl/query"
	"usersvc/domain"
)

func (receiver UserDaoImpl) CheckUsernameExists(ctx context.Context, username string) (bool, error) {
	many, err := receiver.SelectMany(ctx, query.C().Col("username").Eq(username))
	if err != nil {
		return false, err
	}
	users := many.([]domain.User)
	if len(users) > 0 {
		return true, nil
	}
	return false, nil
}
```
这样就实现了对生成的dao层代码的自定义扩展。以后如果user实体的字段新增或者减少，只需要删除userdaosql.go文件，再次执行`go-doudou ddl --dao --pre=t_`命令，重新生成userdaosql.go文件即可，已存在的dao层文件不会被修改。
然后就是SignUp方法的具体实现了
```go
func (receiver *UsersvcImpl) PublicSignUp(ctx context.Context, username string, password string, code *string) (data int, err error) {
	hashPassword, _ := lib.HashPassword(password)
	userDao := dao.NewUserDao(receiver.db)
	var exists bool
	exists, err = userDao.CheckUsernameExists(ctx, username)
	if err != nil {
		panic(err)
	}
	if exists {
		return 0, ddhttp.NewBizError(lib.ErrUsernameExists, ddhttp.WithStatusCode(406))
	}
	user := domain.User{
		Username: username,
		Password: hashPassword,
	}
	_, err = userDao.Insert(ctx, &user)
	if err != nil {
		panic(err)
	}
	return user.Id, nil
}
```
遇到报错，可以直接panic，也可以`return "", lib.ErrUsernameExists`。因为已经加了`ddhttp.Recover`中间件，可以自动从panic里恢复，并返回错误信息给前端。需要注意的是，http状态码为500，不是200。只要从接口方法里返回了error类型的参数，生成的http handler代码里默认设置的http状态码就是500。如果想自定义修改默认生成的http handler里的代码，是完全可以的。当有接口定义新增或者修改的时候，再次执行命令`go-doudou svc http --handler -c go --doc`不会覆盖已存在的代码，只会增量生成代码。
#### Postman测试
测试一下接口，这是第一次请求

![image.png](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/60584cd54667405981635745c2023881~tplv-k3u1fbpfcp-watermark.image?)
这是第二次请求

![image.png](https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/a0dc242cdd4742ee91014235f3425752~tplv-k3u1fbpfcp-watermark.image?)

### 用户登录接口
#### PublicLogIn方法实现
```go
func (receiver *UsersvcImpl) PublicLogIn(ctx context.Context, username string, password string) (data string, err error) {
	userDao := dao.NewUserDao(receiver.db)
	many, err := userDao.SelectMany(ctx, query.C().Col("username").Eq(username).And(query.C().Col("delete_at").IsNull()))
	if err != nil {
		panic(err)
	}
	users := many.([]domain.User)
	if len(users) == 0 || !lib.CheckPasswordHash(password, users[0].Password) {
		return "", ddhttp.NewBizError(lib.ErrUsernameOrPasswordIncorrect, ddhttp.WithStatusCode(401))
	}
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": users[0].Id,
		"exp":    now.Add(12 * time.Hour).Unix(),
		//"iat":    now.Unix(),
		//"nbf":    now.Unix(),
	})
	return token.SignedString([]byte(receiver.conf.BizConf.JwtSecret))
}
```
这段代码的逻辑是先根据入参username查出来数据库中的用户，如果没查到或者密码不对，返回“用户名或密码错误”的报错，如果密码对了，则签发token返回。用的jwt库是[golang-jwt/jwt](https://github.com/golang-jwt/jwt)。

#### Postman测试

![image.png](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/4eb20348dab64907b1c8ae58fb3052a8~tplv-k3u1fbpfcp-watermark.image?)

### 上传头像接口
#### 修改domain
表里面少了一个avatar字段，现在我们加上：
```go
Avatar   string     `dd:"type:varchar(255);extra:comment '用户头像'"`
```
因为是新增了字段，所以要先删除dao/userdaosql.go文件，再执行ddl命令
```shell
go-doudou ddl --dao --pre=t_
```
如果增删的字段比较多，涉及多个实体，可以通过如下命令一次删掉所有`*sql.go`文件，再重新生成
```
rm -rf dao/*sql.go
```
#### 修改.env配置
加入三行配置。JWT_为前缀的是JWT token校验相关的配置。Biz_为前缀的是实际业务相关的配置。
```shell
JWT_SECRET=secret
JWT_IGNORE_URL=/public/sign/up,/public/log/in,/public/get/download/avatar,/public/**
BIZ_OUTPUT=out
```
`JWT_IGNORE_URL`的值设置成`/public/**`就可以了，我都列上，是想说明这里同时支持通配符匹配和完整匹配。
同时，config/config.go文件也需要相应的改动。当然也可以直接调用`os.Getenv`方法。
```go
package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	DbConf  DbConfig
	JWTConf JWTConf
	BizConf BizConf
}

type BizConf struct {
	Output string
}

type JWTConf struct {
	Secret    []byte
	IgnoreUrl []string `split_words:"true"`
}

type DbConfig struct {
	Driver  string `default:"mysql"`
	Host    string `default:"localhost"`
	Port    string `default:"3306"`
	User    string
	Passwd  string
	Schema  string
	Charset string `default:"utf8mb4"`
}

func LoadFromEnv() *Config {
	var dbconf DbConfig
	err := envconfig.Process("db", &dbconf)
	if err != nil {
		logrus.Panicln("Error processing env", err)
	}
	var jwtConf JWTConf
	err = envconfig.Process("jwt", &jwtConf)
	if err != nil {
		logrus.Panicln("Error processing env", err)
	}
	var bizConf BizConf
	err = envconfig.Process("biz", &bizConf)
	if err != nil {
		logrus.Panicln("Error processing env", err)
	}
	return &Config{
		dbconf,
		jwtConf,
		bizConf,
	}
}
```
#### JWT校验中间件
因为go-doudou的http router采用的是gorilla/mux，所以与gorilla/mux的middleware是完全兼容的，自定义中间件的写法也是完全一样的。
```go
package middleware

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gobwas/glob"
	"net/http"
	"os"
	"strings"
)

type ctxKey int

const userIdKey ctxKey = ctxKey(0)

func NewContext(ctx context.Context, id int) context.Context {
	return context.WithValue(ctx, userIdKey, id)
}

func FromContext(ctx context.Context) (int, bool) {
	userId, ok := ctx.Value(userIdKey).(int)
	return userId, ok
}

func Jwt(inner http.Handler) http.Handler {
	g := glob.MustCompile(fmt.Sprintf("{%s}", os.Getenv("JWT_IGNORE_URL")))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if g.Match(r.RequestURI) {
			inner.ServeHTTP(w, r)
			return
		}
		authHeader := r.Header.Get("Authorization")
		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			w.WriteHeader(401)
			w.Write([]byte("Unauthorised.\n"))
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		if userId, exists := claims["userId"]; !exists {
			w.WriteHeader(401)
			w.Write([]byte("Unauthorised.\n"))
			return
		} else {
			inner.ServeHTTP(w, r.WithContext(NewContext(r.Context(), int(userId.(float64)))))
		}
	})
}
```
#### UploadAvatar方法实现
```go
func (receiver *UsersvcImpl) UploadAvatar(ctx context.Context, avatar *v3.FileModel) (data string, err error) {
	defer avatar.Close()
	_ = os.MkdirAll(receiver.conf.BizConf.Output, os.ModePerm)
	out := filepath.Join(receiver.conf.BizConf.Output, avatar.Filename)
	var f *os.File
	f, err = os.OpenFile(out, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = io.Copy(f, avatar.Reader)
	if err != nil {
		panic(err)
	}
	userId, _ := middleware.FromContext(ctx)
	userDao := dao.NewUserDao(receiver.db)
	_, err = userDao.UpdateNoneZero(ctx, domain.User{
		Id:     userId,
		Avatar: out,
	})
	if err != nil {
		panic(err)
	}
	return "OK", nil
}
```
这里需要注意的是，`defer avatar.Close()`这行代码一定要尽早写上，这是释放文件描述符资源的代码。

### 下载头像接口
#### GetPublicDownloadAvatar方法实现
```go
func (receiver *UsersvcImpl) GetPublicDownloadAvatar(ctx context.Context, userId int) (data *os.File, err error) {
	userDao := dao.NewUserDao(receiver.db)
	var get interface{}
	get, err = userDao.Get(ctx, userId)
	if err != nil {
		panic(err)
	}
	return os.Open(get.(domain.User).Avatar)
}
```

### 用户详情接口
#### GetUser方法实现
```go
func (receiver *UsersvcImpl) GetUser(ctx context.Context, userId int) (data vo.UserVo, err error) {
	userDao := dao.NewUserDao(receiver.db)
	var get interface{}
	get, err = userDao.Get(ctx, userId)
	if err != nil {
		panic(err)
	}
	user := get.(domain.User)
	return vo.UserVo{
		Id:       user.Id,
		Username: user.Username,
		Name:     user.Name,
		Phone:    user.Phone,
		Dept:     user.Dept,
	}, nil
}
```
#### Postman测试

![image.png](https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/469925abcde94893985e126f7723d8f5~tplv-k3u1fbpfcp-watermark.image?)

### 用户分页接口
#### 导入测试数据
```sql
INSERT INTO tutorial.t_user (id, username, password, name, phone, dept, create_at, update_at, delete_at, avatar) VALUES (2, 'peter', '$2a$14$VaQLa/GbLAhRZvvTlgE8OOQgsBY4RDAJC5jkz13kjP9RlntdKBZVW', '张三丰', '13552053960', '技术部', '2021-12-28 06:41:00', '2021-12-28 14:59:20', null, 'out/wolf-wolves-snow-wolf-landscape-985ca149f06cd03b9f0ed8dfe326afdb.jpg');
INSERT INTO tutorial.t_user (id, username, password, name, phone, dept, create_at, update_at, delete_at, avatar) VALUES (4, 'john', '$2a$14$AKCs.u9vFUOCe5VwcmdfwOAkeiDtQYEgIB/nSU8/eemYwd91.qU.i', '李世民', '13552053961', '行政部', '2021-12-28 12:12:32', '2021-12-28 14:59:20', null, '');
INSERT INTO tutorial.t_user (id, username, password, name, phone, dept, create_at, update_at, delete_at, avatar) VALUES (5, 'lucy', '$2a$14$n0.l54axUqnKGagylQLu7ee.yDrtLubxzM1qmOaHK9Ft2P09YtQUS', '朱元璋', '13552053962', '销售部', '2021-12-28 12:13:17', '2021-12-28 14:59:20', null, '');
INSERT INTO tutorial.t_user (id, username, password, name, phone, dept, create_at, update_at, delete_at, avatar) VALUES (6, 'jack', '$2a$14$jFCwiZHcD7.DL/teao.Dl.HAFwk8wM2f1riH1fG2f52WYKqSiGZlC', '张无忌', '', '总裁办', '2021-12-28 12:14:19', '2021-12-28 14:59:20', null, '');
```
#### PageUsers方法实现
```go
func (receiver *UsersvcImpl) PageUsers(ctx context.Context, pageQuery vo.PageQuery) (data vo.PageRet, err error) {
	userDao := dao.NewUserDao(receiver.db)
	var q query.Q
	q = query.C().Col("delete_at").IsNull()
	if stringutils.IsNotEmpty(pageQuery.Filter.Name) {
		q = q.And(query.C().Col("name").Like(fmt.Sprintf(`%s%%`, pageQuery.Filter.Name)))
	}
	if stringutils.IsNotEmpty(pageQuery.Filter.Dept) {
		q = q.And(query.C().Col("dept").Eq(pageQuery.Filter.Dept))
	}
	var page query.Page
	if len(pageQuery.Page.Orders) > 0 {
		for _, item := range pageQuery.Page.Orders {
			page = page.Order(query.Order{
				Col:  item.Col,
				Sort: sortenum.Sort(item.Sort),
			})
		}
	}
	if pageQuery.Page.PageNo == 0 {
		pageQuery.Page.PageNo = 1
	}
	page = page.Limit((pageQuery.Page.PageNo-1)*pageQuery.Page.Size, pageQuery.Page.Size)
	var ret query.PageRet
	ret, err = userDao.PageMany(ctx, page, q)
	if err != nil {
		panic(err)
	}
	var items []vo.UserVo
	for _, item := range ret.Items.([]domain.User) {
		var userVo vo.UserVo
		_ = copier.DeepCopy(item, &userVo)
		items = append(items, userVo)
	}
	data = vo.PageRet{
		Items:    items,
		PageNo:   ret.PageNo,
		PageSize: ret.PageSize,
		Total:    ret.Total,
		HasNext:  ret.HasNext,
	}
	return data, nil
}
```
#### Postman测试

![image.png](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/ad393790fe294a0d9ae200513490429a~tplv-k3u1fbpfcp-watermark.image?)

### 服务部署
最后介绍一下docker-compose部署服务
首先修改Dockerfile
```
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
```
然后修改docker-compose.yml
```
version: '3.9'

services:
  db:
    container_name: db
    image: mysql:5.7
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: 1234
    ports:
      - 3306:3306
    volumes:
      - $PWD/my:/etc/mysql/conf.d
      - $PWD/sqlscripts:/docker-entrypoint-initdb.d
    networks:
      - tutorial

  usersvc:
    container_name: usersvc
    build:
      context: .
    environment:
      - GDD_BANNER=off
      - GDD_PORT=6060
      - DB_HOST=db
    expose:
      - "6060"
    ports:
      - "6060:6060"
    networks:
      - tutorial
    depends_on:
      - db

networks:
  tutorial:
    driver: bridge
```
最后执行命令
```shell
docker-compose -f docker-compose.yml up -d
```
如果usersvc容器没有启动成功，可能是因为db容器还没有完全启动，可以多执行几遍上面的命令。

### 总结
到这里，我们达到了全部的学习目标，也实现了需求清单中的全部接口。教程的全部源码都在[这里](https://github.com/unionj-cloud/go-doudou-tutorials/tree/master/usersvc)。








