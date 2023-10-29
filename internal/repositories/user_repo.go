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

func (r *Repository) GetUsersByParameter(parameters string, page int) ([]*model.User, error) {
	users := make([]*model.User, 0)
	countSkipUsers := (page - 1) * 5
	query := "SELECT id, name, surname, patronymic, age, gender, country FROM Users" + parameters + " offset $1 limit 5"

	err := r.db.Select(&users, query, countSkipUsers)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r Repository) UpdateUserById(id int, updateInfo string) (int64, error) {
	query := "UPDATE Users set " + updateInfo + " where id = $1"
	res, err := r.db.Exec(query, id)
	if err != nil {
		return 0, err
	}
	resStatus, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return resStatus, nil
}

func (r Repository) DeleteUserById(id int) error {
	query := "DELETE FROM Users WHERE id = $1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
