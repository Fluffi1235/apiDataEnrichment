package methods

import (
	"Effective_Mobile/internal/services"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func AllUsers(r chi.Router, service *services.Service) {
	r.Get("/allUsers", func(w http.ResponseWriter, r *http.Request) {
		page := r.FormValue("page")
		users, err := service.GetAllUsers(page)
		if err != nil {
			log.Println(err)
			return
		}
		if len(users) == 0 {
			_, err = w.Write([]byte(fmt.Sprintf("Больше пользователей нет, страница %s", page)))
			if err != nil {
				log.Println(err)
				return
			}
		}

		var usersJSN [][]byte
		for _, user := range users {
			userJSN, err := json.Marshal(&user)
			if err != nil {
				log.Println(err)
				usersJSN = append(usersJSN, userJSN)
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
}

func SearchUsersByParameter(r chi.Router, service *services.Service) {
	r.Get("/search", func(w http.ResponseWriter, r *http.Request) {
		parameter := r.FormValue("parameter")
		value := r.FormValue("value")
		page := r.FormValue("page")
		if value == "" || parameter == "" {
			_, err := w.Write([]byte("Неверно введены параметры"))
			if err != nil {
				log.Println(err)
				return
			}
		}

		users, err := service.GetUsersByParameter(parameter, value, page)
		if err != nil {
			log.Println(err)
			return
		}

		if len(users) == 0 {
			_, err = w.Write([]byte(fmt.Sprintf("Больше пользователей нет, страница %s", parameter)))
			if err != nil {
				log.Println(err)
				return
			}
		}

		var usersJSN [][]byte
		for _, user := range users {
			userJSN, err := json.Marshal(&user)
			if err != nil {
				log.Println(err)
			}
			usersJSN = append(usersJSN, userJSN)
		}

		for _, userJSN := range usersJSN {
			_, err = w.Write(userJSN)
			if err != nil {
				log.Println(err)
				return
			}
		}

		log.Printf("Return all users by parameter %s", parameter)
	})
}
