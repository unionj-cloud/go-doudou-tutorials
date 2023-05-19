/**
* Generated by go-doudou v2.0.8.
* Don't edit!
 */
package client

import (
	"context"
	"testgorm/dto"

	"github.com/go-resty/resty/v2"
)

type Options struct {
	GzipReqBody bool
}

type ITestgormClient interface {
	PostCadCommonComment(ctx context.Context, _headers map[string]string, body dto.CadCommonComment, options Options) (_resp *resty.Response, data string, err error)
	GetCadCommonComment_Id(ctx context.Context, _headers map[string]string, id string, options Options) (_resp *resty.Response, data dto.CadCommonComment, err error)
	PutCadCommonComment(ctx context.Context, _headers map[string]string, body dto.CadCommonComment, options Options) (_resp *resty.Response, re error)
	DeleteCadCommonComment_Id(ctx context.Context, _headers map[string]string, id string, options Options) (_resp *resty.Response, re error)
	GetCadCommonComments(ctx context.Context, _headers map[string]string, parameter dto.Parameter, options Options) (_resp *resty.Response, data dto.Page, err error)
	PostCadCommonLike(ctx context.Context, _headers map[string]string, body dto.CadCommonLike, options Options) (_resp *resty.Response, data string, err error)
	GetCadCommonLike_Id(ctx context.Context, _headers map[string]string, id string, options Options) (_resp *resty.Response, data dto.CadCommonLike, err error)
	PutCadCommonLike(ctx context.Context, _headers map[string]string, body dto.CadCommonLike, options Options) (_resp *resty.Response, re error)
	DeleteCadCommonLike_Id(ctx context.Context, _headers map[string]string, id string, options Options) (_resp *resty.Response, re error)
	GetCadCommonLikes(ctx context.Context, _headers map[string]string, parameter dto.Parameter, options Options) (_resp *resty.Response, data dto.Page, err error)
	PostQrtzJobDetail(ctx context.Context, _headers map[string]string, body dto.QrtzJobDetail, options Options) (_resp *resty.Response, data string, err error)
	GetQrtzJobDetail_Id(ctx context.Context, _headers map[string]string, id string, options Options) (_resp *resty.Response, data dto.QrtzJobDetail, err error)
	PutQrtzJobDetail(ctx context.Context, _headers map[string]string, body dto.QrtzJobDetail, options Options) (_resp *resty.Response, re error)
	DeleteQrtzJobDetail_Id(ctx context.Context, _headers map[string]string, id string, options Options) (_resp *resty.Response, re error)
	GetQrtzJobDetails(ctx context.Context, _headers map[string]string, parameter dto.Parameter, options Options) (_resp *resty.Response, data dto.Page, err error)
	PostTInvalidToken(ctx context.Context, _headers map[string]string, body dto.TInvalidToken, options Options) (_resp *resty.Response, data int32, err error)
	GetTInvalidToken_Id(ctx context.Context, _headers map[string]string, id int32, options Options) (_resp *resty.Response, data dto.TInvalidToken, err error)
	PutTInvalidToken(ctx context.Context, _headers map[string]string, body dto.TInvalidToken, options Options) (_resp *resty.Response, re error)
	DeleteTInvalidToken_Id(ctx context.Context, _headers map[string]string, id int32, options Options) (_resp *resty.Response, re error)
	GetTInvalidTokens(ctx context.Context, _headers map[string]string, parameter dto.Parameter, options Options) (_resp *resty.Response, data dto.Page, err error)
	PostCadCommonConfig(ctx context.Context, _headers map[string]string, body dto.CadCommonConfig, options Options) (_resp *resty.Response, data string, err error)
	GetCadCommonConfig_Id(ctx context.Context, _headers map[string]string, id string, options Options) (_resp *resty.Response, data dto.CadCommonConfig, err error)
	PutCadCommonConfig(ctx context.Context, _headers map[string]string, body dto.CadCommonConfig, options Options) (_resp *resty.Response, re error)
	DeleteCadCommonConfig_Id(ctx context.Context, _headers map[string]string, id string, options Options) (_resp *resty.Response, re error)
	GetCadCommonConfigs(ctx context.Context, _headers map[string]string, parameter dto.Parameter, options Options) (_resp *resty.Response, data dto.Page, err error)
	PostQrtzBlobTrigger(ctx context.Context, _headers map[string]string, body dto.QrtzBlobTrigger, options Options) (_resp *resty.Response, data string, err error)
	GetQrtzBlobTrigger_Id(ctx context.Context, _headers map[string]string, id string, options Options) (_resp *resty.Response, data dto.QrtzBlobTrigger, err error)
	PutQrtzBlobTrigger(ctx context.Context, _headers map[string]string, body dto.QrtzBlobTrigger, options Options) (_resp *resty.Response, re error)
	DeleteQrtzBlobTrigger_Id(ctx context.Context, _headers map[string]string, id string, options Options) (_resp *resty.Response, re error)
	GetQrtzBlobTriggers(ctx context.Context, _headers map[string]string, parameter dto.Parameter, options Options) (_resp *resty.Response, data dto.Page, err error)
	PostQrtzCalendar(ctx context.Context, _headers map[string]string, body dto.QrtzCalendar, options Options) (_resp *resty.Response, data string, err error)
	GetQrtzCalendar_Id(ctx context.Context, _headers map[string]string, id string, options Options) (_resp *resty.Response, data dto.QrtzCalendar, err error)
	PutQrtzCalendar(ctx context.Context, _headers map[string]string, body dto.QrtzCalendar, options Options) (_resp *resty.Response, re error)
	DeleteQrtzCalendar_Id(ctx context.Context, _headers map[string]string, id string, options Options) (_resp *resty.Response, re error)
	GetQrtzCalendars(ctx context.Context, _headers map[string]string, parameter dto.Parameter, options Options) (_resp *resty.Response, data dto.Page, err error)
	PostTUser(ctx context.Context, _headers map[string]string, body dto.TUser, options Options) (_resp *resty.Response, data int32, err error)
	GetTUser_Id(ctx context.Context, _headers map[string]string, id int32, options Options) (_resp *resty.Response, data dto.TUser, err error)
	PutTUser(ctx context.Context, _headers map[string]string, body dto.TUser, options Options) (_resp *resty.Response, re error)
	DeleteTUser_Id(ctx context.Context, _headers map[string]string, id int32, options Options) (_resp *resty.Response, re error)
	GetTUsers(ctx context.Context, _headers map[string]string, parameter dto.Parameter, options Options) (_resp *resty.Response, data dto.Page, err error)
	PostTWordCloudTask(ctx context.Context, _headers map[string]string, body dto.TWordCloudTask, options Options) (_resp *resty.Response, data int32, err error)
	GetTWordCloudTask_Id(ctx context.Context, _headers map[string]string, id int32, options Options) (_resp *resty.Response, data dto.TWordCloudTask, err error)
	PutTWordCloudTask(ctx context.Context, _headers map[string]string, body dto.TWordCloudTask, options Options) (_resp *resty.Response, re error)
	DeleteTWordCloudTask_Id(ctx context.Context, _headers map[string]string, id int32, options Options) (_resp *resty.Response, re error)
	GetTWordCloudTasks(ctx context.Context, _headers map[string]string, parameter dto.Parameter, options Options) (_resp *resty.Response, data dto.Page, err error)
	PostCadCommonFile(ctx context.Context, _headers map[string]string, body dto.CadCommonFile, options Options) (_resp *resty.Response, data string, err error)
	GetCadCommonFile_Id(ctx context.Context, _headers map[string]string, id string, options Options) (_resp *resty.Response, data dto.CadCommonFile, err error)
	PutCadCommonFile(ctx context.Context, _headers map[string]string, body dto.CadCommonFile, options Options) (_resp *resty.Response, re error)
	DeleteCadCommonFile_Id(ctx context.Context, _headers map[string]string, id string, options Options) (_resp *resty.Response, re error)
	GetCadCommonFiles(ctx context.Context, _headers map[string]string, parameter dto.Parameter, options Options) (_resp *resty.Response, data dto.Page, err error)
	PostCadCommonFollow(ctx context.Context, _headers map[string]string, body dto.CadCommonFollow, options Options) (_resp *resty.Response, data string, err error)
	GetCadCommonFollow_Id(ctx context.Context, _headers map[string]string, id string, options Options) (_resp *resty.Response, data dto.CadCommonFollow, err error)
	PutCadCommonFollow(ctx context.Context, _headers map[string]string, body dto.CadCommonFollow, options Options) (_resp *resty.Response, re error)
	DeleteCadCommonFollow_Id(ctx context.Context, _headers map[string]string, id string, options Options) (_resp *resty.Response, re error)
	GetCadCommonFollows(ctx context.Context, _headers map[string]string, parameter dto.Parameter, options Options) (_resp *resty.Response, data dto.Page, err error)
	PostQrtzFiredTrigger(ctx context.Context, _headers map[string]string, body dto.QrtzFiredTrigger, options Options) (_resp *resty.Response, data string, err error)
	GetQrtzFiredTrigger_Id(ctx context.Context, _headers map[string]string, id string, options Options) (_resp *resty.Response, data dto.QrtzFiredTrigger, err error)
	PutQrtzFiredTrigger(ctx context.Context, _headers map[string]string, body dto.QrtzFiredTrigger, options Options) (_resp *resty.Response, re error)
	DeleteQrtzFiredTrigger_Id(ctx context.Context, _headers map[string]string, id string, options Options) (_resp *resty.Response, re error)
	GetQrtzFiredTriggers(ctx context.Context, _headers map[string]string, parameter dto.Parameter, options Options) (_resp *resty.Response, data dto.Page, err error)
	PostTClient(ctx context.Context, _headers map[string]string, body dto.TClient, options Options) (_resp *resty.Response, data int32, err error)
	GetTClient_Id(ctx context.Context, _headers map[string]string, id int32, options Options) (_resp *resty.Response, data dto.TClient, err error)
	PutTClient(ctx context.Context, _headers map[string]string, body dto.TClient, options Options) (_resp *resty.Response, re error)
	DeleteTClient_Id(ctx context.Context, _headers map[string]string, id int32, options Options) (_resp *resty.Response, re error)
	GetTClients(ctx context.Context, _headers map[string]string, parameter dto.Parameter, options Options) (_resp *resty.Response, data dto.Page, err error)
	PostQrtzCronTrigger(ctx context.Context, _headers map[string]string, body dto.QrtzCronTrigger, options Options) (_resp *resty.Response, data string, err error)
	GetQrtzCronTrigger_Id(ctx context.Context, _headers map[string]string, id string, options Options) (_resp *resty.Response, data dto.QrtzCronTrigger, err error)
	PutQrtzCronTrigger(ctx context.Context, _headers map[string]string, body dto.QrtzCronTrigger, options Options) (_resp *resty.Response, re error)
	DeleteQrtzCronTrigger_Id(ctx context.Context, _headers map[string]string, id string, options Options) (_resp *resty.Response, re error)
	GetQrtzCronTriggers(ctx context.Context, _headers map[string]string, parameter dto.Parameter, options Options) (_resp *resty.Response, data dto.Page, err error)
}
