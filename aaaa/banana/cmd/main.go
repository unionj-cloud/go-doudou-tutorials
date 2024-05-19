/**
* Generated by go-doudou v2.3.3.
* You can edit it as your need.
 */
package main

import (
	_ "aaaa/banana/plugin"

	"github.com/unionj-cloud/go-doudou/v2/framework/grpcx"
	"github.com/unionj-cloud/go-doudou/v2/framework/plugin"
	"github.com/unionj-cloud/go-doudou/v2/framework/rest"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/zlogger"
)

func main() {
	srv := rest.NewRestServer()
	grpcServer := grpcx.NewEmptyGrpcServer()
	plugins := plugin.GetServicePlugins()
	for _, key := range plugins.Keys() {
		value, _ := plugins.Get(key)
		value.Initialize(srv, grpcServer, nil)
	}
	defer func() {
		if r := recover(); r != nil {
			zlogger.Info().Msgf("Recovered. Error: %v\n", r)
		}
		for _, key := range plugins.Keys() {
			value, _ := plugins.Get(key)
			value.Close()
		}
	}()
	go func() {
		grpcServer.Run()
	}()
	srv.Run()
}
