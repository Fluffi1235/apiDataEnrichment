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

func UpdateInfoUser(r chi.Router, service *services.Service) {
	r.Put("/update", func(w http.ResponseWriter, r *http.Request) {
		filters, err := io.ReadAll(r.Body)
		decoder := json.NewDecoder(bytes.NewReader(filters))
		updatedUser := &model.User{}

		err = decoder.Decode(updatedUser)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}

		if updatedUser.Id <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte("Неверно введен id пользователя"))
			if err != nil {
				log.Println(err)
				return
			}
			return
		}

		log.Printf("Updating user id: %d", updatedUser.Id)

		status, err := service.UpdateUser(updatedUser)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}

		if status == 0 {
			_, err = w.Write([]byte("Пользователя с таким id не существует"))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(err)
				return
			}
			w.WriteHeader(http.StatusOK)
			return
		}

		_, err = w.Write([]byte(fmt.Sprintf("Пользователь с id=%d обновлен", updatedUser.Id)))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}

		log.Printf("User with id=%d updated", updatedUser.Id)
	})
}
