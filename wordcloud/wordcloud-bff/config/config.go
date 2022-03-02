package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"time"
)

type Config struct {
	RedisConf RedisConf
	BizConf   BizConf
	ConConf   ConcurrencyConf
}

type RedisConf struct {
	Host string
}

type ConcurrencyConf struct {
	RatelimitRate       float64       `split_words:"true"`
	RatelimitBurst      int           `split_words:"true"`
	BulkheadWorkers     int           `split_words:"true"`
	BulkheadMaxwaittime time.Duration `split_words:"true"`
}

type BizConf struct {
	Output      string
	OssEndpoint string `split_words:"true"`
	OssKey      string `split_words:"true"`
	OssSecret   string `split_words:"true"`
	OssBucket   string `split_words:"true"`
	JwtToken    string `split_words:"true"`
}

func LoadFromEnv() *Config {
	var redisConf RedisConf
	err := envconfig.Process("redis", &redisConf)
	if err != nil {
		logrus.Panicln("Error processing env", err)
	}
	var bizConf BizConf
	err = envconfig.Process("biz", &bizConf)
	if err != nil {
		logrus.Panicln("Error processing env", err)
	}
	var conConf ConcurrencyConf
	err = envconfig.Process("concurrency", &conConf)
	if err != nil {
		logrus.Panicln("Error processing env", err)
	}
	return &Config{
		redisConf,
		bizConf,
		conConf,
	}
}
