package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	BizConf BizConf
}

type BizConf struct {
	Output      string
	OssEndpoint string `split_words:"true"`
	OssKey      string `split_words:"true"`
	OssSecret   string `split_words:"true"`
	OssBucket   string `split_words:"true"`
}

func LoadFromEnv() *Config {
	var bizConf BizConf
	err := envconfig.Process("biz", &bizConf)
	if err != nil {
		logrus.Panicln("Error processing env", err)
	}
	return &Config{
		bizConf,
	}
}
