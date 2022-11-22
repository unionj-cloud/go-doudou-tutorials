package service

import (
	"context"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/vo"
)

//go:generate go-doudou svc http
//go:generate go-doudou svc grpc

// Usersvc 用户管理服务
// 调用用户详情、用户分页和上传头像接口需要带上Bearer Token请求头
// 用户注册、用户登录和下载头像接口可以公开访问，无须鉴权
// Usersvc is user management service
// You should set Bearer Token header when you request protected endpoints such as user detail, user pagination and upload avatar.
// You can add doc for whole service here
type Usersvc interface {
	// GetUser 用户详情接口
	// 展示如何定义带查询字符串参数的GET请求接口
	// GetUser is user detail api
	// demo how to define get http request with query string parameters
	// @role(USER)
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

	// GetMe 获取当前登录用户详情接口
	// GetMe is used for getting user info from token in header.
	// @role(USER)
	GetMe(ctx context.Context) (
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
	SignUp(ctx context.Context,
		// 用户名
		// username
		username string,
		// 密码
		// password
		password string,
	) (
		// 成功返回用户ID
		// return user ID if success
		data int, err error)

	// PublicLogIn 用户登录接口
	// 展示如何鉴权并返回token
	// PublicLogIn is user login api
	// demo how to do authentication and issue token
	LogIn(ctx context.Context,
		// 用户名
		// username
		username string,
		// 密码
		// password
		password string) (
		// token
		data string, err error)

	// GetLogout 注销token
	// GetLogout is used for revoking a token
	// https://github.com/dgrijalva/jwt-go/issues/214
	// @role(USER)
	GetLogout(ctx context.Context) (
		// 成功返回OK
		// return OK if success
		data string,
		// 错误信息
		// error
		err error)

	// PublicTokenValidate token校验接口
	// 如果token有效，返回用户信息，否则返回空对象
	// PublicTokenValidate validates token string
	// if token is valid, return user information, otherwise, return empty data
	GetTokenValidate(ctx context.Context, token string) (user vo.UserVo, err error)
}
