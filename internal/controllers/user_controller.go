package controllers

import (
	"Effective_Mobile/internal/controllers/methods"
	"Effective_Mobile/internal/model"
	"Effective_Mobile/internal/services"
	"github.com/go-chi/chi"
)

func UserController(r *chi.Mux, cfg *model.Config, service *services.Service) error {
	r.Route("/api", func(r chi.Router) {
		methods.SearchUsersByParameter(r, service)
		methods.CreateUser(r, service, cfg)
		methods.UpdateInfoUser(r, service)
		methods.DeleteUserById(r, service)
	})
	return nil
}
