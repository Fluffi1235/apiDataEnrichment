package methods

import (
	"Effective_Mobile/internal/model"
	"Effective_Mobile/internal/services"
	"encoding/json"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func CreateUser(r chi.Router, service *services.Service, cfg *model.Config) {
	r.Post("/createUser", func(w http.ResponseWriter, r *http.Request) {
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

		user, err := service.CreateUser(name, surname, patronymic, cfg)
		if err != nil {
			log.Println(err)
			return
		}

		userJSN, err := json.Marshal(&user)
		if err != nil {
			log.Println(err)
		}

		_, err = w.Write(userJSN)
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("User created")
	})
}
