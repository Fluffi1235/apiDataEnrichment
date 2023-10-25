package methods

import (
	"Effective_Mobile/internal/services"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func DeleteUserById(r chi.Router, service *services.Service) {
	r.Delete("/deleteUserById", func(w http.ResponseWriter, r *http.Request) {
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
}
