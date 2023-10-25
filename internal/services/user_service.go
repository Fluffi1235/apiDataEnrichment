package services

import (
	"Effective_Mobile/internal/model"
	"Effective_Mobile/internal/parsers"
	"Effective_Mobile/internal/repositories"
	"encoding/json"
)

type Repository struct {
	repo repositories.Repo
}

func New(repo repositories.Repo) *Repository {
	return &Repository{
		repo: repo,
	}
}

func (r *Repository) CreateUser(name, surname, patronymic string, cfg *model.Config) (error, []byte) {
	user := &model.User{-1, name, surname, patronymic, -1, "", ""}
	var err error
	user.Age, err = parsers.GetUserAge(name, cfg)
	if err != nil {
		return err, nil
	}
	user.Gender, err = parsers.GetUserGender(name, cfg)
	if err != nil {
		return err, nil
	}
	user.Country, err = parsers.GetUserCountry(name, cfg)
	if err != nil {
		return err, nil
	}
	err = r.repo.SaveUser(user)
	if err != nil {
		return err, nil
	}
	userJSN, err := json.Marshal(&user)
	if err != nil {
		return err, nil
	}
	return nil, userJSN
}

func (r *Repository) UpdateUser(newUser *model.User) error {
	err := r.repo.UpdateUserById(newUser)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteUserById(id string) error {
	err := r.repo.DeleteUserById(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetAllUsers(page string) (error, [][]byte) {
	users, err := r.repo.GetInfoAllUsers(page)
	if err != nil {
		return err, nil
	}
	var usersJSN [][]byte
	for _, user := range users {
		userJSN, err := json.Marshal(&user)
		if err != nil {
			return err, nil
		}
		usersJSN = append(usersJSN, userJSN)
	}
	return nil, usersJSN
}

func (r *Repository) GetUsersByParameter(parameter, value, page string) (error, [][]byte) {
	var users []*model.User
	var err error
	if isValidParameter(parameter) {
		users, err = r.repo.GetUsersByParameter(page, parameter, value)
		if err != nil {
			return err, nil
		}
	} else {
		return nil, [][]byte{[]byte("Такого параметра не существует")}
	}
	var usersJSN [][]byte
	for _, user := range users {
		userJSN, err := json.Marshal(&user)
		if err != nil {
			return err, nil
		}
		usersJSN = append(usersJSN, userJSN)
	}
	return nil, usersJSN
}

func isValidParameter(parameter string) bool {
	if parameter == "name" || parameter == "surname" || parameter == "patronymic" ||
		parameter == "age" || parameter == "gender" || parameter == "country" {
		return true
	}
	return false
}
