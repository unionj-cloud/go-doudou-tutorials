package service

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
	"usersvc/config"
	"usersvc/dao"
	"usersvc/domain"
	"usersvc/internal/lib"
	"usersvc/internal/middleware"
	"usersvc/vo"

	"github.com/unionj-cloud/go-doudou/toolkit/sqlext/wrapper"

	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/unionj-cloud/go-doudou/toolkit/copier"
	"github.com/unionj-cloud/go-doudou/toolkit/sqlext/query"
	"github.com/unionj-cloud/go-doudou/toolkit/sqlext/sortenum"
	"github.com/unionj-cloud/go-doudou/toolkit/stringutils"

	v3 "github.com/unionj-cloud/go-doudou/toolkit/openapi/v3"
)

type UsersvcImpl struct {
	conf *config.Config
	db   wrapper.DB
}

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
func (receiver *UsersvcImpl) GetUser(ctx context.Context, userId int) (data vo.UserVo, err error) {
	userDao := dao.NewUserDao(receiver.db)
	var get interface{}
	get, err = userDao.Get(ctx, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return vo.UserVo{}, ddhttp.NewBizError(lib.ErrUserNotFound, ddhttp.WithStatusCode(406))
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

func NewUsersvc(conf *config.Config, db wrapper.DB) Usersvc {
	return &UsersvcImpl{
		conf,
		db,
	}
}
