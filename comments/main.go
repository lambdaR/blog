package main

import (
	"github.com/micro/blog/comments/handler"
	pb "github.com/micro/blog/comments/proto"
	"go-micro.dev/v5"
)

func main() {
	service := micro.NewService(
		micro.Name("comments"),
	)

	pb.RegisterCommentsHandler(service.Server(), handler.New())

	service.Init()

	service.Run()
}
