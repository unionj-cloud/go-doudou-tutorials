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
