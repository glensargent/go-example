package main

import "github.com/glensargent/go-example/service"

func main() {
	// var client *redis.Client
	svc := service.New()
	svc.MakeCache()
	svc.MakeRouter()
	svc.Start()
}
