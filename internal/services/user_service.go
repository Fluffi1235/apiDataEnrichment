package services

import (
	"Effective_Mobile/internal/model"
	"Effective_Mobile/internal/parsers"
	"Effective_Mobile/internal/repositories"
	"errors"
	"strconv"
)

type Service struct {
	repo repositories.Repo
}

func New(repo repositories.Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateUser(name, surname, patronymic string, cfg *model.Config) (*model.User, error) {
	user := &model.User{-1, name, surname, patronymic, -1, "", ""}
	var err error

	user.Age, err = parsers.GetUserAge(name, cfg)
	if err != nil {
		return nil, err
	}

	user.Gender, err = parsers.GetUserGender(name, cfg)
	if err != nil {
		return nil, err
	}

	user.Country, err = parsers.GetUserCountry(name, cfg)
	if err != nil {
		return nil, err
	}

	err = s.repo.SaveUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) UpdateUser(newUser *model.User) error {
	err := s.repo.UpdateUserById(newUser)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteUserById(id string) error {
	idUser, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = s.repo.DeleteUserById(idUser)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetAllUsers(page string) ([]*model.User, error) {
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}
	users, err := s.repo.GetInfoAllUsers(pageInt)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Service) GetUsersByParameter(parameter, value, page string) ([]*model.User, error) {
	var users []*model.User
	var err error

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}

	if isValidParameter(parameter) {
		users, err = s.repo.GetUsersByParameter(pageInt, parameter, value)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Неверно введен фильтр")
	}

	return users, nil
}

func isValidParameter(parameter string) bool {
	if parameter == "name" || parameter == "surname" || parameter == "patronymic" ||
		parameter == "age" || parameter == "gender" || parameter == "country" {
		return true
	}

	return false
}
