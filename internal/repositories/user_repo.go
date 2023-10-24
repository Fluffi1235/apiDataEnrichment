package repositories

import (
	"Effective_Mobile/internal/model"
	"log"
	"strconv"
)

func (r Repository) SaveUser(user *model.User) error {
	query := "INSERT INTO Users(name, surname, patronymic, age, gender, country) VALUES($1, $2, $3, $4, $5, $6) RETURNING id"
	err := r.db.Get(&user.Id, query, user.Name, user.SurName, user.Patronymic, user.Age, user.Gender, user.Country)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) GetInfoAllUsers(page string) ([]*model.User, error) {
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}
	countSkipUsers := (pageInt - 1) * 5
	users := make([]*model.User, 0)
	query := "SELECT id, name, surname, patronymic, age, gender, country FROM Users offset $1 limit 5"
	err = r.db.Select(&users, query, countSkipUsers)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r Repository) UpdateUserById(newUser *model.User) error {
	query := "UPDATE Users set name = $1, surname = $2, patronymic = $3, age = $4, gender = $5, country = $6 where id = $7"
	_, err := r.db.Exec(query, newUser.Name, newUser.SurName, newUser.Patronymic, newUser.Age, newUser.Gender, newUser.Country, newUser.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) DeleteUserById(id string) error {
	idUser, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	query := "DELETE FROM Users WHERE id = $1"
	_, err = r.db.Exec(query,
		idUser)
	if err != nil {
		return err
	}
	log.Printf("User with id = %s deleted", id)
	return nil
}
