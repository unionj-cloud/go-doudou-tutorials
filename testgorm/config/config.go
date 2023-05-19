/**
* Generated by go-doudou v2.0.8.
* You can edit it as your need.
*/
package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/errorx"
)

type Config struct {
	BizConf BizConf
}

type BizConf struct {
	ApiSecret string `split_words:"true"`
}

func LoadFromEnv() *Config {
	var bizConf BizConf
	err := envconfig.Process("biz", &bizConf)
	if err != nil {
		errorx.Panic("Error processing environment variables")
	}
	return &Config{
		BizConf: bizConf,
	}
}
