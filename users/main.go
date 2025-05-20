package main

import (
	"go-micro.dev/v5"

	"github.com/micro/blog/users/handler"
	pb "github.com/micro/blog/users/proto"
)

func main() {
	service := micro.New("users")

	pb.RegisterUsersHandler(service.Server(), handler.New())

	service.Init()

	service.Run()
}
