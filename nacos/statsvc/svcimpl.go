package service

import (
	"context"
	"statsvc/config"
	"statsvc/internal/nacosservicej"
)

type StatsvcImpl struct {
	conf       *config.Config
	echoClient *nacosservicej.EchoClient
}

func (receiver *StatsvcImpl) Add(ctx context.Context, x int, y int) (data int, err error) {
	return x + y, nil
}

func NewStatsvc(conf *config.Config, echoClient *nacosservicej.EchoClient) Statsvc {
	return &StatsvcImpl{
		conf,
		echoClient,
	}
}

func (receiver *StatsvcImpl) GetEcho(ctx context.Context, s string) (data string, err error) {
	data, _, err = receiver.echoClient.GetEchoString(ctx, nil, s)
	return
}
