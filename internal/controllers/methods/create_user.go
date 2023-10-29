package methods

import (
	"Effective_Mobile/internal/model"
	"Effective_Mobile/internal/services"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"io"
	"log"
	"net/http"
)

func CreateUser(r chi.Router, service *services.Service, cfg *model.Config) {
	r.Post("/createUser", func(w http.ResponseWriter, r *http.Request) {
		filters, err := io.ReadAll(r.Body)
		decoder := json.NewDecoder(bytes.NewReader(filters))
		user := &model.User{}

		err = decoder.Decode(user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}
		if user.Name == "" || user.SurName == "" {
			_, err = w.Write([]byte("Некорректно введенные данные"))
			if err != nil {
				log.Println(err)
			}
			return
		}

		log.Printf("Creating user %v", user)

		newUser, err := service.CreateUser(user, cfg)
		if err != nil {
			log.Println(err)
			if errors.Is(err, errors.New("Country error")) {
				_, err = w.Write([]byte("Пользователь не зарегистрирован ни в одной стране"))
				if err != nil {
					log.Println(err)
				}
				return
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		userJSN, err := json.Marshal(&newUser)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}

		_, err = w.Write(userJSN)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		log.Printf("User created %v", newUser)
	})
}
