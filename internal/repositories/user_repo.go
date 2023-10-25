package repositories

import (
	"Effective_Mobile/internal/model"
)

func (r Repository) SaveUser(user *model.User) error {
	query := "INSERT INTO Users(name, surname, patronymic, age, gender, country) VALUES($1, $2, $3, $4, $5, $6) RETURNING id"
	err := r.db.Get(&user.Id, query, user.Name, user.SurName, user.Patronymic, user.Age, user.Gender, user.Country)
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) GetInfoAllUsers(page int) ([]*model.User, error) {
	users := make([]*model.User, 0)
	countSkipUsers := (page - 1) * 5
	query := "SELECT id, name, surname, patronymic, age, gender, country FROM Users offset $1 limit 5"

	err := r.db.Select(&users, query, countSkipUsers)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) GetUsersByParameter(page int, parameter, value string) ([]*model.User, error) {
	users := make([]*model.User, 0)
	countSkipUsers := (page - 1) * 5
	query := "SELECT id, name, surname, patronymic, age, gender, country FROM Users where " + parameter + " = $2 offset $3 limit 5"

	err := r.db.Select(&users, query, parameter, value, countSkipUsers)
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

func (r Repository) DeleteUserById(id int) error {
	query := "DELETE FROM Users WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
