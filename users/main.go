package main

import "go-micro.dev/v5"

func main() {
	service := micro.New("users")

	service.Init()

	service.Run()
}
