package config

import (
	"github.com/kelseyhightower/envconfig"
	logger "github.com/unionj-cloud/go-doudou/toolkit/zlogger"
)

type Config struct {
	DbConf  DbConfig
	DepConf DepConfig
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

type DepConfig struct {
	ServerAddr string `split_words:"true"`
}

func LoadFromEnv() *Config {
	var dbconf DbConfig
	err := envconfig.Process("db", &dbconf)
	if err != nil {
		logger.Panic().Err(err).Msg("Error processing env")
	}
	var depConf DepConfig
	err = envconfig.Process("dep", &depConf)
	if err != nil {
		logger.Panic().Err(err).Msg("Error processing env")
	}
	return &Config{
		DbConf:  dbconf,
		DepConf: depConf,
	}
}
