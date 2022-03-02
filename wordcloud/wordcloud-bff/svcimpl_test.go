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
	ddhttp "github.com/unionj-cloud/go-doudou/framework/http"
	"github.com/unionj-cloud/go-doudou/toolkit/sqlext/sortenum"
	"testing"
)

var svc service.WordcloudBff

func TestMain(m *testing.M) {
	conf := config.LoadFromEnv()
	userRestyClient := ddhttp.NewClient()
	userRestyClient.SetHeader("Authorization", fmt.Sprintf("Bearer %s", conf.BizConf.JwtToken))
	userClient := userclient.NewUsersvcClient(ddhttp.WithClient(userRestyClient))
	taskClient := taskclient.NewWordcloudTaskClient()
	rec := metrics.NewPrometheusRecorder(prometheus.DefaultRegisterer)
	userClientProxy := userclient.NewUsersvcClientProxy(userClient, rec)
	taskClientProxy := taskclient.NewWordcloudTaskClientProxy(taskClient, rec)
	svc = service.NewWordcloudBff(conf, nil, nil, taskClientProxy, userClientProxy)

	m.Run()
}

func TestWordcloudBffImpl_TaskPage(t *testing.T) {
	ctx := service.NewUserIdContext(context.Background(), 1)
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
