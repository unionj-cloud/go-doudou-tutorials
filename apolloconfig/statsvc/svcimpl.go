package service

import (
	"context"
	"statsvc/config"
)

type StatsvcImpl struct {
	conf *config.Config
}

func (receiver *StatsvcImpl) Add(ctx context.Context, x int, y int) (data int, err error) {
	return x + y, nil
}

func NewStatsvc(conf *config.Config) Statsvc {
	return &StatsvcImpl{
		conf,
	}
}
