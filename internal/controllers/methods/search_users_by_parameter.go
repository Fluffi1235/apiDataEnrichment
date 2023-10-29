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
			w.WriteHeader(http.StatusBadRequest)
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
			w.WriteHeader(http.StatusOK)
			_, err = w.Write([]byte(fmt.Sprintf("Пользователей нет, страница %d", parameters.Page)))
			if err != nil {
				log.Println(err)
				return
			}
		}
		usersJSN, err := json.Marshal(&users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		_, err = w.Write(usersJSN)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		log.Printf("Return all users by parameter %v", parameters)
	})
}
