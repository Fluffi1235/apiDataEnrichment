package controllers

import (
	"Effective_Mobile/internal/model"
	"Effective_Mobile/internal/services"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"strconv"
)

func UserController(r *chi.Mux, cfg *model.Config, service *services.Repository) error {
	r.Route("/api", func(r chi.Router) {
		r.Get("/allUsers", func(w http.ResponseWriter, r *http.Request) {
			page := r.FormValue("page")
			err, usersJSN := service.GetAllUsers(page)
			if err != nil {
				log.Println(err)
				return
			}
			if len(usersJSN) == 0 {
				_, err = w.Write([]byte(fmt.Sprintf("Больше пользователей нет, страница %s", page)))
				if err != nil {
					log.Println(err)
					return
				}
			}
			for _, userJSN := range usersJSN {
				_, err = w.Write(userJSN)
				if err != nil {
					log.Println(err)
					return
				}
			}
			log.Printf("Return all users page %s", page)
		})
		r.Get("/createUser", func(w http.ResponseWriter, r *http.Request) {
			if r.FormValue("name") == "" || r.FormValue("surname") == "" {
				_, err := w.Write([]byte("Некорректно введенные данные"))
				if err != nil {
					log.Println(err)
					return
				}
			}
			name := r.FormValue("name")
			surname := r.FormValue("surname")
			patronymic := r.FormValue("patronymic")
			log.Printf("Creating user name: %s surname: %s patronymic: %s", name, surname, patronymic)
			err, userJSN := service.CreateUser(name, surname, patronymic, cfg)
			if err != nil {
				log.Println(err)
				return
			}
			_, err = w.Write(userJSN)
			if err != nil {
				log.Println(err)
				return
			}
			log.Printf("User created")
		})
		r.Get("/update", func(w http.ResponseWriter, r *http.Request) {
			id := r.FormValue("id")
			age := r.FormValue("age")
			if id == "" {
				_, err := w.Write([]byte("Не ввели id пользователя"))
				if err != nil {
					log.Println(err)
					return
				}
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				log.Println(err)
				return
			}
			ageInt, err := strconv.Atoi(age)
			if err != nil {
				log.Println(err)
				return
			}
			newUser := &model.User{
				Id:         idInt,
				Name:       r.FormValue("name"),
				SurName:    r.FormValue("surname"),
				Patronymic: r.FormValue("patronymic"),
				Age:        ageInt,
				Gender:     r.FormValue("gender"),
				Country:    r.FormValue("country"),
			}
			log.Printf("Updating user id: %s", id)
			err = service.UpdateUser(newUser)
			if err != nil {
				log.Println(err)
				return
			}
			_, err = w.Write([]byte(fmt.Sprintf("Пользователь с id=%s обновлен %v", id, newUser)))
			if err != nil {
				log.Println(err)
				return
			}
			log.Printf("User with id=%s updated %v", id, newUser)
		})
		r.Get("/deleteUserById", func(w http.ResponseWriter, r *http.Request) {
			id := r.FormValue("id")
			if id == "" {
				_, err := w.Write([]byte("Не ввели id пользователя"))
				if err != nil {
					log.Println(err)
					return
				}
			}
			log.Printf("Deleting user with id=%s", id)
			err := service.DeleteUserById(id)
			if err != nil {
				log.Println(err)
				return
			}
			_, err = w.Write([]byte(fmt.Sprintf("Пользователь с id=%s удален", id)))
			if err != nil {
				log.Println(err)
				return
			}
			log.Printf("User with id=%s deleted", id)
		})
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte("pong"))
			if err != nil {
				log.Println(err)
				return
			}
		})
	})
	return nil
}
