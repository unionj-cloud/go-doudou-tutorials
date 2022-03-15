package service_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	service "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/config"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/dao"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/db"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/domain"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/vo"
	"github.com/unionj-cloud/go-doudou/toolkit/sqlext/sortenum"
	"testing"
	"time"
)

var svc service.WordcloudTask
var taskDao dao.WordCloudTaskDao
var conn *sqlx.DB

func TestMain(m *testing.M) {
	conf := config.LoadFromEnv()
	var err error
	conn, err = db.NewDb(conf.DbConf)
	if err != nil {
		panic(err)
	}
	defer func() {
		if conn == nil {
			return
		}
		if err := conn.Close(); err == nil {
			logrus.Infoln("Database connection is closed")
		} else {
			logrus.Warnln("Failed to close database connection")
		}
	}()

	svc = service.NewWordcloudTask(conf, conn)

	taskDao = dao.NewWordCloudTaskDao(conn)
	for i := 0; i < 25; i++ {
		_, err = taskDao.Insert(context.Background(), domain.WordCloudTask{
			Id:     0,
			SrcUrl: "abc" + fmt.Sprint(i),
			ImgUrl: "abc",
			Lang:   "zh",
			Status: 2,
			Error:  "",
			UserId: 1,
		})
		if err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}

	m.Run()

	//got, err := taskDao.DeleteMany(context.Background(), query.C().Col("delete_at").IsNull())
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(got)
}

func TestWordcloudTaskImpl_TaskPage(t *testing.T) {
	gotData, err := svc.TaskPage(context.Background(), vo.PageQuery{
		Filter: vo.Filter{
			UserId: 1,
		},
		Page: vo.Page{
			Orders: []vo.Order{
				{
					Col:  "create_at",
					Sort: string(sortenum.Desc),
				},
			},
			PageNo: 1,
			Size:   10,
		},
	})
	require.NoError(t, err)
	data, _ := json.MarshalIndent(gotData, "", "  ")
	fmt.Println(string(data))
}
