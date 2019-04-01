package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"miner/srv/user/handler"
	"miner/srv/user/subscriber"

	example "miner/srv/user/proto/example"
	user "miner/srv/user/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("ars.miner.srv.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("ars.miner.srv.user", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("ars.miner.srv.user", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
