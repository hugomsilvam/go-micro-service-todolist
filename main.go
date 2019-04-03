package main

import (
	"github.com/hugomsilvam/todo.list.micro/handler"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	example "github.com/hugomsilvam/todo.list.micro/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.todo.list.micro"),
		micro.Version("0.1"))

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterTodoServiceHandler(service.Server(), new(handler.Todo))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
