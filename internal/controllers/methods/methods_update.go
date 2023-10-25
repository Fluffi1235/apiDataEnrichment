package methods

import (
	"Effective_Mobile/internal/model"
	"Effective_Mobile/internal/services"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"strconv"
)

func UpdateInfoUser(r chi.Router, service *services.Service) {
	r.Put("/update", func(w http.ResponseWriter, r *http.Request) {
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
}
