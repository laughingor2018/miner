package main

import (
	"github.com/micro/go-log"

	"github.com/micro/go-micro"
	"github.com/laughingor2018/miner/api/handler"
	"github.com/laughingor2018/miner/api/client"

	example "github.com/laughingor2018/miner/api/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("ars.miner.api.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the Example srv client
		micro.WrapHandler(client.ExampleWrapper(service)),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}