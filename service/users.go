package service

import (
	"net/http"

	"github.com/go-chi/render"
)

type response struct {
	Status int    `json:"status"`
	String string `json:"string"`
}

func (svc *Service) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, response{
		String: "helloo",
	})
}
