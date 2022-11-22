package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/zlogger"
)

type Config struct {
	DbConf  DbConfig
	BizConf BizConf
}

type DbConfig struct {
	Driver  string `default:"mysql"`
	Host    string `default:"localhost"`
	Port    string `default:"3306"`
	User    string
	Passwd  string
	Schema  string
	Charset string `default:"utf8mb4"`
}

type BizConf struct {
	Output      string
	OssEndpoint string `split_words:"true"`
	OssKey      string `split_words:"true"`
	OssSecret   string `split_words:"true"`
	OssBucket   string `split_words:"true"`
}

func LoadFromEnv() *Config {
	var dbconf DbConfig
	err := envconfig.Process("db", &dbconf)
	if err != nil {
		zlogger.Panic().Err(err).Msg("Error processing env")
	}
	var bizConf BizConf
	err = envconfig.Process("biz", &bizConf)
	if err != nil {
		zlogger.Panic().Err(err).Msg("Error processing env")
	}
	return &Config{
		DbConf:  dbconf,
		BizConf: bizConf,
	}
}
