package methods

import (
	"Effective_Mobile/internal/model"
	"Effective_Mobile/internal/services"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"io"
	"log"
	"net/http"
)

func SearchUsersByParameter(r chi.Router, service *services.Service) {
	r.Post("/search", func(w http.ResponseWriter, r *http.Request) {
		filters, err := io.ReadAll(r.Body)
		decoder := json.NewDecoder(bytes.NewReader(filters))
		parameters := &model.Filter{}
		err = decoder.Decode(parameters)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		if parameters.Page < 0 {
			_, err = w.Write([]byte("Введите страницу %s"))
			if err != nil {
				log.Println(err)
				return
			}
		}
		users, err := service.GetUsersByParameter(parameters)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		if len(users) == 0 {
			_, err = w.Write([]byte(fmt.Sprintf("Пользователей нет, страница %d", parameters.Page)))
			if err != nil {
				log.Println(err)
				return
			}
		}

		var usersJSN [][]byte
		for _, user := range users {
			userJSN, err := json.Marshal(&user)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(err)
			}
			usersJSN = append(usersJSN, userJSN)
		}

		for _, userJSN := range usersJSN {
			_, err = w.Write(userJSN)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(err)
				return
			}
		}

		log.Printf("Return all users by parameter %v", parameters)
	})
}
