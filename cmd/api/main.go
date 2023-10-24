package main

import (
	"Effective_Mobile/internal/config"
	"Effective_Mobile/internal/controllers"
	"Effective_Mobile/internal/repositories"
	"Effective_Mobile/internal/services"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.ReadConfigEnv()
	if err != nil {
		log.Fatal(err)
	}
	db, err := sqlx.Connect("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatal(err)
	}
	repo := repositories.New(db)
	service := services.New(repo)
	defer db.Close()
	r := chi.NewRouter()
	err = controllers.UserController(r, cfg, service)
	if err != nil {
		log.Println(err)
	}
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
