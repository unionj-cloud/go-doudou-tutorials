package service

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/sqlext/wrapper"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/config"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/dao"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/domain"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/internal/lib"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/vo"

	"github.com/pkg/errors"
	pb "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/transport/grpc"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/copier"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/sqlext/query"
)

var _ pb.UsersvcServiceServer = (*UsersvcImpl)(nil)

var _ Usersvc = (*UsersvcImpl)(nil)

type UsersvcImpl struct {
	pb.UnimplementedUsersvcServiceServer

	conf *config.Config
	db   wrapper.GddDB
}

type ctxKey int

const userIdKey ctxKey = ctxKey(0)
const tokenKey ctxKey = ctxKey(1)

func NewUserIdContext(ctx context.Context, id int) context.Context {
	return context.WithValue(ctx, userIdKey, id)
}

func UserIdFromContext(ctx context.Context) (int, bool) {
	userId, ok := ctx.Value(userIdKey).(int)
	return userId, ok
}

func NewTokenContext(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenKey, token)
}

func TokenFromContext(ctx context.Context) (string, bool) {
	token, ok := ctx.Value(tokenKey).(string)
	return token, ok
}

func (receiver *UsersvcImpl) GetUser(ctx context.Context, userId int) (data vo.UserVo, err error) {
	userDao := dao.NewUserDao(receiver.db)
	var user domain.User
	err = userDao.Get(ctx, &user, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return vo.UserVo{}, rest.NewBizError(lib.ErrUserNotFound, rest.WithStatusCode(406))
		} else {
			panic(err)
		}
	}
	copier.DeepCopy(user, &data)
	return
}
func (receiver *UsersvcImpl) SignUp(ctx context.Context, username string, password string) (data int, err error) {
	hashPassword, _ := lib.HashPassword(password)
	userDao := dao.NewUserDao(receiver.db)
	var exists bool
	exists, err = userDao.CheckUsernameExists(ctx, username)
	if err != nil {
		panic(err)
	}
	if exists {
		return 0, rest.NewBizError(lib.ErrUsernameExists, rest.WithStatusCode(406))
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
func (receiver *UsersvcImpl) LogIn(ctx context.Context, username string, password string) (data string, err error) {
	userDao := dao.NewUserDao(receiver.db)
	var users []domain.User
	err = userDao.SelectMany(ctx, &users, query.C().Col("username").Eq(username).ToWhere())
	if err != nil {
		panic(err)
	}
	if len(users) == 0 || !lib.CheckPasswordHash(password, users[0].Password) {
		return "", rest.NewBizError(lib.ErrUsernameOrPasswordIncorrect, rest.WithStatusCode(401))
	}
	return lib.JwtToken(receiver.conf.BizConf.JwtSecret, jwt.MapClaims{
		"userId": users[0].Id,
		"exp":    time.Now().Add(12 * time.Hour).Unix(),
		//"iat":    now.Unix(),
		//"nbf":    now.Unix(),
	})
}

func (receiver *UsersvcImpl) GetMe(ctx context.Context) (data vo.UserVo, err error) {
	userId, _ := UserIdFromContext(ctx)
	userDao := dao.NewUserDao(receiver.db)
	var user domain.User
	err = userDao.Get(ctx, &user, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return vo.UserVo{}, rest.NewBizError(lib.ErrUserNotFound, rest.WithStatusCode(406))
		}
		panic(err)
	}
	copier.DeepCopy(user, &data)
	return
}

func (receiver *UsersvcImpl) GetLogout(ctx context.Context) (data string, err error) {
	tokenString, _ := TokenFromContext(ctx)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("BIZ_JWT_SECRET")), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	expireAt := time.Unix(int64(claims["exp"].(float64)), 0).Local()
	tokenDao := dao.NewInvalidTokenDao(receiver.db)
	_, err = tokenDao.UpsertNoneZero(ctx, &domain.InvalidToken{
		Token:    tokenString,
		ExpireAt: &expireAt,
	})
	if err != nil {
		panic(err)
	}
	return "OK", nil
}

func ValidateToken(ctx context.Context, tokenString string, db wrapper.GddDB) (int, bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("BIZ_JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		return 0, false, nil
	}

	claims := token.Claims.(jwt.MapClaims)
	if userId, exists := claims["userId"]; !exists {
		return 0, false, nil
	} else {
		tokenDao := dao.NewInvalidTokenDao(db)
		tctx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()
		cnt, err := tokenDao.CountMany(tctx, query.C().Col("token").Eq(tokenString).ToWhere())
		if err != nil {
			return 0, false, err
		}
		if cnt > 0 {
			return 0, false, nil
		}
		return int(userId.(float64)), true, nil
	}
}

func (receiver *UsersvcImpl) GetTokenValidate(ctx context.Context, token string) (user vo.UserVo, err error) {
	userId, valid, err := ValidateToken(ctx, token, receiver.db)
	if err != nil || !valid {
		return vo.UserVo{}, err
	}
	return receiver.GetUser(ctx, userId)
}

func (receiver *UsersvcImpl) GetUserRpc(ctx context.Context, request *pb.GetUserRpcRequest) (*pb.UserVo, error) {
	data, err := receiver.GetUser(ctx, int(request.UserId))
	if err != nil {
		return nil, err
	}
	var ret pb.UserVo
	copier.DeepCopy(data, &ret)
	return &ret, nil
}
func (receiver *UsersvcImpl) GetMeRpc(ctx context.Context, request *emptypb.Empty) (*pb.UserVo, error) {
	data, err := receiver.GetMe(ctx)
	if err != nil {
		return nil, err
	}
	var ret pb.UserVo
	copier.DeepCopy(data, &ret)
	return &ret, nil
}
func (receiver *UsersvcImpl) SignUpRpc(ctx context.Context, request *pb.SignUpRpcRequest) (*pb.SignUpRpcResponse, error) {
	data, err := receiver.SignUp(ctx, request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	return &pb.SignUpRpcResponse{
		Data: int32(data),
	}, nil
}
func (receiver *UsersvcImpl) LogInRpc(ctx context.Context, request *pb.LogInRpcRequest) (*pb.LogInRpcResponse, error) {
	data, err := receiver.LogIn(ctx, request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	return &pb.LogInRpcResponse{
		Data: data,
	}, nil
}
func (receiver *UsersvcImpl) GetLogoutRpc(ctx context.Context, request *emptypb.Empty) (*pb.GetLogoutRpcResponse, error) {
	data, err := receiver.GetLogout(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetLogoutRpcResponse{
		Data: data,
	}, nil
}
func (receiver *UsersvcImpl) GetTokenValidateRpc(ctx context.Context, request *pb.GetTokenValidateRpcRequest) (*pb.UserVo, error) {
	data, err := receiver.GetTokenValidate(ctx, request.Token)
	if err != nil {
		return nil, err
	}
	var ret pb.UserVo
	copier.DeepCopy(data, &ret)
	return &ret, nil
}

func NewUsersvc(conf *config.Config, db wrapper.GddDB) *UsersvcImpl {
	return &UsersvcImpl{
		conf: conf,
		db:   db,
	}
}
