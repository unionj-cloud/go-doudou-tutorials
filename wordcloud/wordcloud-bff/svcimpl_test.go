package service_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	service "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff"
	"github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-bff/config"
	taskpb "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-task/transport/grpc"
	userclient "github.com/unionj-cloud/go-doudou-tutorials/wordcloud/wordcloud-user/client"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/zlogger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"testing"
	"time"
)

var svc service.WordcloudBff

func TestMain(m *testing.M) {
	conf := config.LoadFromEnv()
	userClient := userclient.NewUsersvcClient()

	tlsOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	dialOptions := []grpc.DialOption{
		tlsOption,
	}

	serverAddr := os.Getenv("WORDCLOUDTASK")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	taskClientConn, err := grpc.DialContext(ctx, serverAddr, dialOptions...)
	if err != nil {
		zlogger.Panic().Err(err).Msgf("[go-doudou] failed to connect to server %s", serverAddr)
	}
	defer taskClientConn.Close()

	svc = service.NewWordcloudBff(conf, nil, nil, taskpb.NewWordcloudTaskServiceClient(taskClientConn), userClient)

	m.Run()
}

func TestWordcloudBffImpl_TaskPage(t *testing.T) {
	testToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGllbnRJZCI6IndvcmRjbG91ZC1iZmYiLCJ1c2VySWQiOjB9.Do7TxRRT_jyUmBjcNR2dKmkTFNwW5iEZFDBVYdJTVQE"
	ctx := service.NewUserIdContext(service.NewTokenContext(context.Background(), testToken), 1)
	page, err := svc.GetTaskPage(ctx, 1, 5)
	require.NoError(t, err)
	data, _ := json.MarshalIndent(page, "", "  ")
	fmt.Println(string(data))
}
