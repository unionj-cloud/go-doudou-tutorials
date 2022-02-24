# go-doudou series 01: How to develop a monolithic RESTful service with go-doudou

[EN](./README.md) [中文](./README_zh.md)

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
### TOC

- [TODO](#todo)
- [Learning Goals](#learning-goals)
- [Dev Requirements](#dev-requirements)
- [Install go-doudou](#install-go-doudou)
- [Init Project](#init-project)
- [Define Apis](#define-apis)
- [Generate Code](#generate-code)
- [Start Service](#start-service)
- [Init Database](#init-database)
- [Generate Domain and CRUD Code](#generate-domain-and-crud-code)
- [User SignUp](#user-signup)
  - [Fix Domain](#fix-domain)
  - [PublicSignUp Method](#publicsignup-method)
  - [Test by Postman](#test-by-postman)
- [User LogIn](#user-login)
  - [PublicLogIn Method](#publiclogin-method)
  - [Test by Postman](#test-by-postman-1)
- [Upload Avatar](#upload-avatar)
  - [Fix Domain](#fix-domain-1)
  - [Fix .env](#fix-env)
  - [JWT Middleware](#jwt-middleware)
  - [UploadAvatar Method](#uploadavatar-method)
- [Download Avatar](#download-avatar)
  - [GetPublicDownloadAvatar Method](#getpublicdownloadavatar-method)
- [User Pagination](#user-pagination)
  - [Import Test Data](#import-test-data)
  - [PageUsers Method](#pageusers-method)
  - [Test by Postman](#test-by-postman-2)
- [User Detail](#user-detail)
  - [GetUser Method](#getuser-method)
  - [Test by Postman](#test-by-postman-3)
- [Deploy Service](#deploy-service)
- [Source Code](#source-code)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

In this tutorial, I will introduce [go-doudou microservice framework](https://github.com/unionj-cloud/go-doudou) to you. Go-doudou has built-in service register, discover and fault tolerance features based on gossip protocol, and it uses gorilla/mux as http router, and uses golang interface as [IDL(Interface Definition Language)](https://en.wikipedia.org/wiki/Interface_description_language). Go-doudou is an IDL compiler and server/client code generator tool at first, then it becomes a microservice framework. Its RESTful version is stable and production ready, while grpc version is in early development.

After brief introduction, now I will use a user management service as an example to demonstrate how to develop a monolithic service with go-doudou. This will be the first post about go-doudou, we will begin from monolithic and then go to microservice step by step.

&nbsp;
### TODO
- User SignUp
- User LogIn
- User Detail
- User Pagination
- Upload Avatar
- Download Avatar

&nbsp;
### Learning Goals
- Protect user detail, user pagination and upload avatar endpoints by jwt token
- Online OpenAPI 3.0 api documentation
- Golang http client code
- Fake response for frontend
- Use built-in ddl tool for syncing structs to tables and generating structs from tables
- Use generated dao layer code for single table CRUD

&nbsp;
### Dev Requirements
- docker: recommend to install Desktop from [official website](https://docs.docker.com/engine/install/)
- IDE：goland

&nbsp;
### Install go-doudou
- If go version is below 1.16：
```shell
GO111MODULE=on  go get -v github.com/unionj-cloud/go-doudou@v0.9.8
```
if go version is 1.16 or above：
```shell
go get -v github.com/unionj-cloud/go-doudou@v0.9.8
```
- Try to execute command `go-doudou -v`, if it printed below content, installation is successful：
```shell
➜  ~ go-doudou -v
go-doudou version v0.9.8
```
if it printed command not found, you should try to check `GOPATH`, and the go-doudou binary should be in `$GOPATH/bin` directory.
```shell
➜  ~ go env | grep GOPATH
GOPATH="/Users/wubin1989/go"
```
I can find go-doudou binary in `/Users/wubin1989/go/bin` directory for example. Then you should add one line `export PATH=/Users/wubin1989/go/bin:$PATH` to .zshrc or .bashrc file. And then execute `source ~/.zshrc` or `source ~/.bashrc` or open a new command line tab to make it work.

&nbsp;
### Init Project
Execute command in arbitrary directory. Here is in tutorials directory：
```
go-doudou svc init usersvc 
```
cd usersvc, you will see below structure：
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
- svc.go file：IDL file for defining api endpoints by interface methods
- vo folder：define structs as api input and output parameters. Vo is named from view object.
- Dockerfile：for building docker image

&nbsp;
### Define Apis
Have a look at svc.go file：
```go
package service

import (
	"context"
	v3 "github.com/unionj-cloud/go-doudou/openapi/v3"
	"os"
	"usersvc/vo"
)

// Usersvc is user management service
// You should set Bearer Token header when you request protected endpoints such as user detail, user pagination and upload avatar.
// You can add doc for whole service here
type Usersvc interface {
	// PageUsers is user pagination api
	// demo how to define post request api which accepts application/json content-type
	PageUsers(ctx context.Context,
		// pagination parameter
		query vo.PageQuery) (
		// pagination result
		data vo.PageRet,
		// error
		err error)

	// GetUser is user detail api
	// demo how to define get http request with query string parameters
	GetUser(ctx context.Context,
		// user id
		userId int) (
		// user detail
		data vo.UserVo,
		// error
		err error)

	// PublicSignUp is user signup api
	// demo how to define post request api which accepts application/x-www-form-urlencoded content-type
	PublicSignUp(ctx context.Context,
		// username
		username string,
		// password
		password string,
		// image code
		code string,
	) (
		// return OK if success
		data string, err error)

	// PublicLogIn is user login api
	// demo how to do authentication and issue token
	PublicLogIn(ctx context.Context,
		// username
		username string,
		// password
		password string) (
		// token
		data string, err error)

	// UploadAvatar is avatar upload api
	// demo how to define file upload api
	// NOTE: there must be at least one []v3.FileModel or v3.FileModel input parameter
	UploadAvatar(ctx context.Context,
		// user avatar
		avatar v3.FileModel) (
		// return OK if success
		data string, err error)

	// GetPublicDownloadAvatar is avatar download api
	// demo how to define file download api
	// NOTE: there must be one and at most one *os.File output parameter
	GetPublicDownloadAvatar(ctx context.Context,
		// user id
		userId int) (
		// avatar file
		data *os.File, err error)
}
```
All comments in above code will be displayed in online OpenAPI 3.0 documentation, so recommend to add comments to let your clients use your apis more easily. For simplicity, no support for fancy symbol like `@`, only support golang `//` syntax.

&nbsp;
### Generate Code
Execute below command to generate server and client code:
```
go-doudou svc http --handler -c go --doc
```
Explain：
- `--handler`: to generate http handler implementation
- `-c`：to generate client http request code. Currently only support go.
- `--doc`：to generate OpenAPI 3.0 doc file in json format
  This is the most used command for me. Recommend to all go-doudou users. When you made some changes in svc.go file, execute this command, and you will get incrementally generated new code. The rule is:
- handler.go file, *_openapi3.json file and client.go file will always be fully overwritten. So you should not edit these files manually.
- handlerimpl.go file, svcimpl.go file and clientproxy.go file will always be incrementally appended new code. Existing code will not be changed by go-doudou. So you can edit these files to fit your needs and still get all benefits from go-doudou.
- Other files will be skipped if already exist.

You'd better execute below command to make sure all dependencies have been downloaded:
```
go mod tidy
```
Let's see project structure now:
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
- client folder: generated go client code
- cmd folder: main.go file in it for starting the service
- config folder: used for populating configs
- db folder: used for connecting database
- svcimpl.go file: there are mock implementations initially to return fake response data for clients. You should replace them with your own implementation later.
- transport folder：there are http handlers in it.
- usersvc_openapi3.go file：used for online doc
- usersvc_openapi3.json file：used for online doc

&nbsp;
### Start Service
```
go-doudou svc run
```
You will see:
```
➜  usersvc git:(master) ✗ go-doudou svc run
INFO[2022-01-23 15:55:07] Initializing logging reporter                
INFO[2022-01-23 15:55:07] ================ Registered Routes ================ 
INFO[2022-01-23 15:55:07] +----------------------+--------+-------------------------+ 
INFO[2022-01-23 15:55:07] |         NAME         | METHOD |         PATTERN         | 
INFO[2022-01-23 15:55:07] +----------------------+--------+-------------------------+ 
INFO[2022-01-23 15:55:07] | PageUsers            | POST   | /page/users             | 
INFO[2022-01-23 15:55:07] | User                 | GET    | /user                   | 
INFO[2022-01-23 15:55:07] | PublicSignUp         | POST   | /public/sign/up         | 
INFO[2022-01-23 15:55:07] | PublicLogIn          | POST   | /public/log/in          | 
INFO[2022-01-23 15:55:07] | UploadAvatar         | POST   | /upload/avatar          | 
INFO[2022-01-23 15:55:07] | PublicDownloadAvatar | GET    | /public/download/avatar | 
INFO[2022-01-23 15:55:07] | GetDoc               | GET    | /go-doudou/doc          | 
INFO[2022-01-23 15:55:07] | GetOpenAPI           | GET    | /go-doudou/openapi.json | 
INFO[2022-01-23 15:55:07] | Prometheus           | GET    | /go-doudou/prometheus   | 
INFO[2022-01-23 15:55:07] | GetRegistry          | GET    | /go-doudou/registry     | 
INFO[2022-01-23 15:55:07] +----------------------+--------+-------------------------+ 
INFO[2022-01-23 15:55:07] =================================================== 
INFO[2022-01-23 15:55:07] Started in 233.424µs                        
INFO[2022-01-23 15:55:07] Http server is listening on :6060      
```
When you see Http server is listening on :6060, it means service has been started and we also have a mock server. For example, we can send a request to `/user` api, to see what will be sent back(I use [httpie](https://github.com/httpie/httpie)):
```
➜  usersvc git:(master) ✗ http http://localhost:6060/user
HTTP/1.1 200 OK
Content-Encoding: gzip
Content-Length: 109
Content-Type: application/json; charset=UTF-8
Date: Sun, 23 Jan 2022 07:59:28 GMT
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
You may notice that all attribute names are capitalised, and that is not what you want. There is one line `go generate` command in vo.go file:
```
//go:generate go-doudou name --file $GOFILE
```
The command uses a built-in small tool called name. It can add or modify json tag to every exported field of each struct in a file. Default mode is camelcase, and it also supports snakecase.
```
go generate ./...
```
Then restart the service, resend a request to `/user` api, you can see it works.
```
➜  usersvc git:(master) ✗ http http://localhost:6060/user
HTTP/1.1 200 OK
Content-Encoding: gzip
Content-Length: 114
Content-Type: application/json; charset=UTF-8
Date: Sun, 23 Jan 2022 08:00:40 GMT
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
If you want to know more about name command, please refer to [README.md](https://github.com/unionj-cloud/go-doudou/blob/main/name/README.md).
At the moment, as json tag of each exported field has been changed, you should regenerate OpenAPI 3.0 doc json file:
```
go-doudou svc http --doc
```
Then restart the service, open browser and navigate to `http://localhost:6060/go-doudou/doc`, input http basic username `admin`, password `admin` to see online doc.

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/7bm15h8385dd86kzn91q.png)


![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/iwulhorfkhko2a350hd0.png)

&nbsp;
### Init Database
For supporting utf8/utf8mb4, you should create `my/custom.cnf` file in project root path and paste below content:
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
Then create `sqlscripts/init.sql` file for initialising database, paste below sqls：
```sql
CREATE SCHEMA `tutorial` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

CREATE TABLE `tutorial`.`t_user`
(
    `id`        INT         NOT NULL AUTO_INCREMENT,
    `username`  VARCHAR(45) NOT NULL COMMENT 'username',
    `password`  VARCHAR(60) NOT NULL COMMENT 'password',
    `name`      VARCHAR(45) NOT NULL COMMENT 'real name',
    `phone`     VARCHAR(45) NOT NULL COMMENT 'phone number',
    `dept`      VARCHAR(45) NOT NULL COMMENT 'department',
    `create_at` DATETIME    NULL DEFAULT current_timestamp,
    `update_at` DATETIME    NULL DEFAULT current_timestamp on update current_timestamp,
    `delete_at` DATETIME    NULL,
    PRIMARY KEY (`id`)
);
```
Create `docker-compose.yml` file, paste below content:
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
Execute `docker compose` command to start mysql docker container:
```
docker-compose -f docker-compose.yml up -d
```
Execute `docker ps` to make sure it started.
```
➜  usersvc git:(master) ✗ docker ps        
CONTAINER ID   IMAGE       COMMAND                  CREATED          STATUS          PORTS                                                  NAMES
df6af6362c41   mysql:5.7   "docker-entrypoint.s…"   13 minutes ago   Up 13 minutes   0.0.0.0:3306->3306/tcp, :::3306->3306/tcp, 33060/tcp   db
```

&nbsp;
### Generate Domain and CRUD Code
As our mysql schema name is tutorial, so we should change the value of environment variable DB_SCHEMA to tutorial in .env file.
```
DB_SCHEMA=tutorial
```
Execute below command to generate domain and dao layer(CRUD) code:
```
go-doudou ddl -r --dao --pre=t_
```
Flag explain：
- `-r`：short for reverse, generate structs from tables
- `--dao`：generate dao layer code, only support single table CRUD operations based on [sqlx](https://github.com/jmoiron/sqlx)
- `--pre`：table name has prefix t_
  At the moment, you can see two more folders:

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/nw8p2697lplkuyi9dgmj.png)

More about the built-in ddl tool, please refer to [README.md](https://github.com/unionj-cloud/go-doudou/blob/main/ddl/doc/README.md)
Here we should see what CRUD operation methods in `dao/base.go` file, we will use them to implement our business logic.
```go
package dao

import (
	"context"
	"github.com/unionj-cloud/go-doudou/ddl/query"
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
We will add `*sqlx.DB` instance as field to `UsersvcImpl` struct in `svcimpl.go` file.
```go
type UsersvcImpl struct {
	conf *config.Config
	db   *sqlx.DB
}
```
Then make a small fix in factory function `NewUsersvc`.
```go
func NewUsersvc(conf *config.Config, db *sqlx.DB) Usersvc {
	return &UsersvcImpl{
		conf,
		db,
	}
}
```
In our main function, it has already injected `*sqlx.DB` instance `conn` into `NewUsersvc`.
```go
svc := service.NewUsersvc(conf, conn)
```
Afterwards, we can use `db` field of `UsersvcImpl` for CRUD operations.

&nbsp;
### User SignUp
#### Fix Domain
Normally username should be unique, so we should fix `domain/user.go` to add unique constraint to `User` struct
```go
Username string     `dd:"type:varchar(45);extra:comment 'username';unique"`
```
Here we can see we add `;unique` to `dd` tag.
Then execute `go-doudou ddl` command:
```shell
go-doudou ddl --pre=t_
```
There is no `-r` flag because we need to sync `User` struct to `t_user` table.

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/f5y5cvfwdd1gh7179dcl.png)

#### PublicSignUp Method
We should add `CheckUsernameExists` method to `UserDao` interface for checking if the username has already been used.
```go
package dao

import "context"

type UserDao interface {
	Base
	CheckUsernameExists(ctx context.Context, username string) (bool, error)
}
```
Then we create a new file `dao/userdaoimplext.go` to add extra method implementations.
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
In this way, we can fulfill more complex needs by adding custom CRUD operation methods to generated dao layer interface. If you add or delete fields of a domain struct, you just need to remove `*daosql.go` file, and generate new one by executing command `go-doudou ddl --dao --pre=t_`, existing dao layer files won't be changed anything.

Then we implement `PublicSignUp` method.
```go
func (receiver *UsersvcImpl) PublicSignUp(ctx context.Context, username string, password string, code string) (data string, err error) {
	hashPassword, _ := lib.HashPassword(password)
	userDao := dao.NewUserDao(receiver.db)
	var exists bool
	exists, err = userDao.CheckUsernameExists(ctx, username)
	if err != nil {
		panic(err)
	}
	if exists {
		panic(lib.ErrUsernameExists)
	}
	_, err = userDao.Insert(ctx, domain.User{
		Username: username,
		Password: hashPassword,
	})
	if err != nil {
		panic(err)
	}
	return "OK", nil
}
```
If error ocurred, you can simply let program panic, or `return "", lib.ErrUsernameExists` because we add a recover middleware `ddhttp.Recover`. It can let our service recover and send error message to clients. If any non-nil error returned, http status code is 500 by default. If you don't like it, you can change generated http handler code manually. Remember existing code in `handlerimpl.go` file won't be changed when execute go-doudou command, only be incrementally appended new code.

#### Test by Postman

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/dr8v27pnta9ivbsgggnj.png)

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/dd9lvhwnmoepv50hu9xz.png)

&nbsp;
### User LogIn
#### PublicLogIn Method
```go
func (receiver *UsersvcImpl) PublicLogIn(ctx context.Context, username string, password string) (data string, err error) {
	userDao := dao.NewUserDao(receiver.db)
	many, err := userDao.SelectMany(ctx, query.C().Col("username").Eq(username).And(query.C().Col("delete_at").IsNull()))
	if err != nil {
		return "", err
	}
	users := many.([]domain.User)
	if len(users) == 0 || !lib.CheckPasswordHash(password, users[0].Password) {
		panic(lib.ErrUsernameOrPasswordIncorrect)
	}
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": users[0].Id,
		"exp":    now.Add(10 * time.Minute).Unix(),
		//"iat":    now.Unix(),
		//"nbf":    now.Unix(),
	})
	return token.SignedString(receiver.conf.JWTConf.Secret)
}
```
The code logic is query user record from database by input parameter `username`, if not found, return `Incorrect username or password` error, if password was correct, issue token. The jwt library used here is [golang-jwt/jwt](https://github.com/golang-jwt/jwt)

#### Test by Postman

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/a3bf5ceit5qusp38kiu8.png)

&nbsp;
### Upload Avatar
#### Fix Domain
We add `avatar` field to `user` struct now:
```go
Avatar   string     `dd:"type:varchar(255);extra:comment 'user avatar'"`
```
Before continue, we should remove `dao/userdaosql.go` file first, then execute ddl command:
```shell
go-doudou ddl --dao --pre=t_
```
If multiple domain structs were changed, we can use wildcard to delete all `*sql.go` files in order to generate new ones.
```
rm -rf dao/*sql.go
```
#### Fix .env
Add three lines. `JWT_` prefixed ones are jwt related configs, while `Biz_` prefixed ones are business related configs.
```shell
JWT_SECRET=secret
JWT_IGNORE_URL=/public/sign/up,/public/log/in,/public/get/download/avatar,/public/**
BIZ_OUTPUT=out
```
You can simply set `JWT_IGNORE_URL` to `/public/**` only. I add so many values here to show what kind of values can be used.

At the same time, `config/config.go` should be changed, too. Of course, you can directly use `os.Getenv`.
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
#### JWT Middleware
As go-doudou is relying on gorilla/mux, so if you can write middlewares for gorilla/mux, you can also write ones for go-doudou.
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

func Jwt(g glob.Glob) func(inner http.Handler) http.Handler {
	return func(inner http.Handler) http.Handler {
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
				return []byte(os.Getenv("BIZ_JWT_SECRET")), nil
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
}
```
Then you can use it like this `srv.AddMiddleware(middleware.Jwt(glob.MustCompile(fmt.Sprintf("{%s}", conf.BizConf.JwtIgnoreUrl))))`.

#### UploadAvatar Method
```go
func (receiver *UsersvcImpl) UploadAvatar(ctx context.Context, avatar v3.FileModel, id int) (data string, err error) {
	defer avatar.Close()
	span := opentracing.SpanFromContext(ctx)
	logrus.Debugln(span)
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
NOTE: don't forget to add `defer avatar.Close()` to release resource.

&nbsp;
### Download Avatar
#### GetPublicDownloadAvatar Method
```go
var (
	errAvatarNotFound = ddhttp.NewBizError(errors.New("avatar not found"))
)

func (receiver *UsersvcImpl) GetPublicDownloadAvatar(ctx context.Context, userId int) (data *os.File, err error) {
	userDao := dao.NewUserDao(receiver.db)
	var get interface{}
	get, err = userDao.Get(ctx, userId)
	if err != nil {
		panic(err)
	}
	avatar := get.(domain.User).Avatar
	if stringutils.IsEmpty(avatar) {
		return nil, errAvatarNotFound
	}
	return os.Open(avatar)
}
```

&nbsp;
### User Pagination
#### Import Test Data
```sql
INSERT INTO tutorial.t_user (id, username, password, name, phone, dept, create_at, update_at, delete_at, avatar) VALUES (2, 'crazyboy', '$2a$14$VaQLa/GbLAhRZvvTlgE8OOQgsBY4RDAJC5jkz13kjP9RlntdKBZVW', 'John Snow', '13552053960', 'IT dept.', '2021-12-28 06:41:00', '2021-12-28 14:59:20', null, 'out/wolf-wolves-snow-wolf-landscape-985ca149f06cd03b9f0ed8dfe326afdb.jpg');
INSERT INTO tutorial.t_user (id, username, password, name, phone, dept, create_at, update_at, delete_at, avatar) VALUES (4, 'david', '$2a$14$AKCs.u9vFUOCe5VwcmdfwOAkeiDtQYEgIB/nSU8/eemYwd91.qU.i', 'David Li', '13552053961', 'Admin dept.', '2021-12-28 12:12:32', '2021-12-28 14:59:20', null, '');
INSERT INTO tutorial.t_user (id, username, password, name, phone, dept, create_at, update_at, delete_at, avatar) VALUES (5, 'lucy', '$2a$14$n0.l54axUqnKGagylQLu7ee.yDrtLubxzM1qmOaHK9Ft2P09YtQUS', 'Lucy Zhang', '13552053962', 'Sales dept.', '2021-12-28 12:13:17', '2021-12-28 14:59:20', null, '');
INSERT INTO tutorial.t_user (id, username, password, name, phone, dept, create_at, update_at, delete_at, avatar) VALUES (6, 'jack', '$2a$14$jFCwiZHcD7.DL/teao.Dl.HAFwk8wM2f1riH1fG2f52WYKqSiGZlC', 'Jack Chen', '', 'CEO Office', '2021-12-28 12:14:19', '2021-12-28 14:59:20', null, '');
```

#### PageUsers Method
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

#### Test by Postman

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/hzxd4259wdobl5vhyxwq.png)

&nbsp;
### User Detail
#### GetUser Method
```go
func (receiver *UsersvcImpl) GetUser(ctx context.Context, userId int) (data vo.UserVo, err error) {
	userDao := dao.NewUserDao(receiver.db)
	var get interface{}
	get, err = userDao.Get(ctx, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return vo.UserVo{}, lib.ErrUserNotFound
		} else {
			panic(err)
		}
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

#### Test by Postman
Before test, you should add token to postman.

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/qhpi3tefy2m8p73rer2r.png)

Then send a request with fake user id.

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/7ni0xhe2ek5zd65s1cd8.png)

Then send a request with real user id.

![Image description](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/78f8cpfolq8v30jd7kw3.png)

&nbsp;
### Deploy Service
I want to show you how to deploy our service by docker-compose.
`Dockerfile`
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
`docker-compose.yml` file
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
Then execute below command:
```shell
docker-compose -f docker-compose.yml up -d
```

&nbsp;
### Source Code
Here is full source code with postman collection file [Click Me](https://github.com/unionj-cloud/go-doudou-tutorials/tree/master/usersvc).
