package service_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/slok/goresilience/metrics"
	"github.com/stretchr/testify/require"
	service "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff/config"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff/vo"
	taskclient "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/client"
	userclient "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/client"
	"github.com/unionj-cloud/go-doudou/toolkit/sqlext/sortenum"
	"testing"
)

var svc service.WordcloudBff

func TestMain(m *testing.M) {
	conf := config.LoadFromEnv()
	userClient := userclient.NewUsersvcClient()
	taskClient := taskclient.NewWordcloudTaskClient()
	rec := metrics.NewPrometheusRecorder(prometheus.DefaultRegisterer)
	userClientProxy := userclient.NewUsersvcClientProxy(userClient, rec)
	taskClientProxy := taskclient.NewWordcloudTaskClientProxy(taskClient, rec)
	svc = service.NewWordcloudBff(conf, nil, nil, taskClientProxy, userClientProxy)

	m.Run()
}

func TestWordcloudBffImpl_TaskPage(t *testing.T) {
	testToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGllbnRJZCI6IndvcmRjbG91ZC1iZmYiLCJ1c2VySWQiOjB9.Do7TxRRT_jyUmBjcNR2dKmkTFNwW5iEZFDBVYdJTVQE"
	ctx := service.NewUserIdContext(service.NewTokenContext(context.Background(), testToken), 1)
	page, err := svc.TaskPage(ctx, vo.PageQuery{
		Page: vo.Page{
			Orders: []vo.Order{
				{
					Col:  "createAt",
					Sort: string(sortenum.Desc),
				},
			},
			PageNo: 1,
			Size:   5,
		},
	})
	require.NoError(t, err)
	data, _ := json.MarshalIndent(page, "", "  ")
	fmt.Println(string(data))
}
