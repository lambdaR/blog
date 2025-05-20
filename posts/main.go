package main

import (
	"github.com/micro/blog/posts/handler"
	pb "github.com/micro/blog/posts/proto"
	"go-micro.dev/v5"
)

func main() {
	service := micro.NewService(
		micro.Name("posts"),
	)

	pb.RegisterPostsHandler(service.Server(), handler.New())

	service.Init()

	service.Run()
}
