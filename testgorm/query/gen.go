// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                = new(Query)
	CadCommonComment *cadCommonComment
	CadCommonConfig  *cadCommonConfig
	CadCommonFile    *cadCommonFile
	CadCommonFollow  *cadCommonFollow
	CadCommonLike    *cadCommonLike
	QrtzBlobTrigger  *qrtzBlobTrigger
	QrtzCalendar     *qrtzCalendar
	QrtzCronTrigger  *qrtzCronTrigger
	QrtzFiredTrigger *qrtzFiredTrigger
	QrtzJobDetail    *qrtzJobDetail
	TClient          *tClient
	TInvalidToken    *tInvalidToken
	TUser            *tUser
	TWordCloudTask   *tWordCloudTask
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	CadCommonComment = &Q.CadCommonComment
	CadCommonConfig = &Q.CadCommonConfig
	CadCommonFile = &Q.CadCommonFile
	CadCommonFollow = &Q.CadCommonFollow
	CadCommonLike = &Q.CadCommonLike
	QrtzBlobTrigger = &Q.QrtzBlobTrigger
	QrtzCalendar = &Q.QrtzCalendar
	QrtzCronTrigger = &Q.QrtzCronTrigger
	QrtzFiredTrigger = &Q.QrtzFiredTrigger
	QrtzJobDetail = &Q.QrtzJobDetail
	TClient = &Q.TClient
	TInvalidToken = &Q.TInvalidToken
	TUser = &Q.TUser
	TWordCloudTask = &Q.TWordCloudTask
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:               db,
		CadCommonComment: newCadCommonComment(db, opts...),
		CadCommonConfig:  newCadCommonConfig(db, opts...),
		CadCommonFile:    newCadCommonFile(db, opts...),
		CadCommonFollow:  newCadCommonFollow(db, opts...),
		CadCommonLike:    newCadCommonLike(db, opts...),
		QrtzBlobTrigger:  newQrtzBlobTrigger(db, opts...),
		QrtzCalendar:     newQrtzCalendar(db, opts...),
		QrtzCronTrigger:  newQrtzCronTrigger(db, opts...),
		QrtzFiredTrigger: newQrtzFiredTrigger(db, opts...),
		QrtzJobDetail:    newQrtzJobDetail(db, opts...),
		TClient:          newTClient(db, opts...),
		TInvalidToken:    newTInvalidToken(db, opts...),
		TUser:            newTUser(db, opts...),
		TWordCloudTask:   newTWordCloudTask(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	CadCommonComment cadCommonComment
	CadCommonConfig  cadCommonConfig
	CadCommonFile    cadCommonFile
	CadCommonFollow  cadCommonFollow
	CadCommonLike    cadCommonLike
	QrtzBlobTrigger  qrtzBlobTrigger
	QrtzCalendar     qrtzCalendar
	QrtzCronTrigger  qrtzCronTrigger
	QrtzFiredTrigger qrtzFiredTrigger
	QrtzJobDetail    qrtzJobDetail
	TClient          tClient
	TInvalidToken    tInvalidToken
	TUser            tUser
	TWordCloudTask   tWordCloudTask
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:               db,
		CadCommonComment: q.CadCommonComment.clone(db),
		CadCommonConfig:  q.CadCommonConfig.clone(db),
		CadCommonFile:    q.CadCommonFile.clone(db),
		CadCommonFollow:  q.CadCommonFollow.clone(db),
		CadCommonLike:    q.CadCommonLike.clone(db),
		QrtzBlobTrigger:  q.QrtzBlobTrigger.clone(db),
		QrtzCalendar:     q.QrtzCalendar.clone(db),
		QrtzCronTrigger:  q.QrtzCronTrigger.clone(db),
		QrtzFiredTrigger: q.QrtzFiredTrigger.clone(db),
		QrtzJobDetail:    q.QrtzJobDetail.clone(db),
		TClient:          q.TClient.clone(db),
		TInvalidToken:    q.TInvalidToken.clone(db),
		TUser:            q.TUser.clone(db),
		TWordCloudTask:   q.TWordCloudTask.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:               db,
		CadCommonComment: q.CadCommonComment.replaceDB(db),
		CadCommonConfig:  q.CadCommonConfig.replaceDB(db),
		CadCommonFile:    q.CadCommonFile.replaceDB(db),
		CadCommonFollow:  q.CadCommonFollow.replaceDB(db),
		CadCommonLike:    q.CadCommonLike.replaceDB(db),
		QrtzBlobTrigger:  q.QrtzBlobTrigger.replaceDB(db),
		QrtzCalendar:     q.QrtzCalendar.replaceDB(db),
		QrtzCronTrigger:  q.QrtzCronTrigger.replaceDB(db),
		QrtzFiredTrigger: q.QrtzFiredTrigger.replaceDB(db),
		QrtzJobDetail:    q.QrtzJobDetail.replaceDB(db),
		TClient:          q.TClient.replaceDB(db),
		TInvalidToken:    q.TInvalidToken.replaceDB(db),
		TUser:            q.TUser.replaceDB(db),
		TWordCloudTask:   q.TWordCloudTask.replaceDB(db),
	}
}

type queryCtx struct {
	CadCommonComment ICadCommonCommentDo
	CadCommonConfig  ICadCommonConfigDo
	CadCommonFile    ICadCommonFileDo
	CadCommonFollow  ICadCommonFollowDo
	CadCommonLike    ICadCommonLikeDo
	QrtzBlobTrigger  IQrtzBlobTriggerDo
	QrtzCalendar     IQrtzCalendarDo
	QrtzCronTrigger  IQrtzCronTriggerDo
	QrtzFiredTrigger IQrtzFiredTriggerDo
	QrtzJobDetail    IQrtzJobDetailDo
	TClient          ITClientDo
	TInvalidToken    ITInvalidTokenDo
	TUser            ITUserDo
	TWordCloudTask   ITWordCloudTaskDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		CadCommonComment: q.CadCommonComment.WithContext(ctx),
		CadCommonConfig:  q.CadCommonConfig.WithContext(ctx),
		CadCommonFile:    q.CadCommonFile.WithContext(ctx),
		CadCommonFollow:  q.CadCommonFollow.WithContext(ctx),
		CadCommonLike:    q.CadCommonLike.WithContext(ctx),
		QrtzBlobTrigger:  q.QrtzBlobTrigger.WithContext(ctx),
		QrtzCalendar:     q.QrtzCalendar.WithContext(ctx),
		QrtzCronTrigger:  q.QrtzCronTrigger.WithContext(ctx),
		QrtzFiredTrigger: q.QrtzFiredTrigger.WithContext(ctx),
		QrtzJobDetail:    q.QrtzJobDetail.WithContext(ctx),
		TClient:          q.TClient.WithContext(ctx),
		TInvalidToken:    q.TInvalidToken.WithContext(ctx),
		TUser:            q.TUser.WithContext(ctx),
		TWordCloudTask:   q.TWordCloudTask.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
