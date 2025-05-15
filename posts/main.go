package main

import "go-micro.dev/v5"

func main() {
	service := micro.New("posts")

	service.Init()

	service.Run()
}
