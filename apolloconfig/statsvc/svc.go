package service

import (
	"context"
)

type Statsvc interface {
	Add(ctx context.Context, x, y int) (data int, err error)
}
