package service

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-redis/redis/v8"
)

type Service struct {
	Cache  *redis.Client
	Router *chi.Mux
}

func (svc *Service) MakeCache() {
	svc.Cache = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
	})

	ctx := context.Background()
	err := svc.Cache.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := svc.Cache.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}

func (svc *Service) MakeRouter() {
	svc.Router = chi.NewRouter()
	svc.Router.Use(middleware.Logger)
	svc.Router.Get("/", svc.HandleGetUsers)
}

func (svc *Service) Start() {
	log.Println("Service starting on 8080")
	http.ListenAndServe(":8080", svc.Router)
}

func New() Service {
	return Service{}
}
