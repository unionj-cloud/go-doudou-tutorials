package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/zlogger"
)

type Config struct {
	DbConf  DbConfig
	BizConf BizConf
}

type BizConf struct {
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
		dbconf,
		bizConf,
	}
}
