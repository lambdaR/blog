package main

import (
	"go-micro.dev/v5"

	"github.com/micro/blog/comments/handler"
	pb "github.com/micro/blog/comments/proto"
)

func main() {
	service := micro.New("comments")

	pb.RegisterCommentsHandler(service.Server(), handler.New())

	service.Init()

	service.Run()
}
