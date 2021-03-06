package client

import (
	"context"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/vo"
	v3 "github.com/unionj-cloud/go-doudou/toolkit/openapi/v3"
)

type IUsersvcClient interface {
	PageUsers(ctx context.Context, _headers map[string]string, query vo.PageQuery) (_resp *resty.Response, data vo.PageRet, err error)
	GetUser(ctx context.Context, _headers map[string]string, userId int) (_resp *resty.Response, data vo.UserVo, err error)
	GetMe(ctx context.Context, _headers map[string]string) (_resp *resty.Response, data vo.UserVo, err error)
	PublicSignUp(ctx context.Context, _headers map[string]string, username string, password string, code *string) (_resp *resty.Response, data int, err error)
	PublicLogIn(ctx context.Context, _headers map[string]string, username string, password string) (_resp *resty.Response, data string, err error)
	UploadAvatar(ctx context.Context, _headers map[string]string, avatar v3.FileModel, id int) (_resp *resty.Response, data string, err error)
	GetPublicDownloadAvatar(ctx context.Context, _headers map[string]string, userId int) (_resp *resty.Response, data *os.File, err error)
	GetLogout(ctx context.Context, _headers map[string]string) (_resp *resty.Response, data string, err error)
	PublicTokenValidate(ctx context.Context, _headers map[string]string, token string) (_resp *resty.Response, user vo.UserVo, err error)
}
