package client

import (
	"context"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	example "github.com/laughingor2018/miner/srv/user/proto/example"
)

type exampleKey struct {}

// FromContext retrieves the client from the Context
func ExampleFromContext(ctx context.Context) (example.ExampleService, bool) {
	c, ok := ctx.Value(exampleKey{}).(example.ExampleService)
	return c, ok
}

// Client returns a wrapper for the ExampleClient
func ExampleWrapper(service micro.Service) server.HandlerWrapper {
	client := example.NewExampleService("ars.miner.srv.user", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, exampleKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}