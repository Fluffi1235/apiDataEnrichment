package services

import (
	"Effective_Mobile/internal/model"
	"Effective_Mobile/internal/parsers"
	"Effective_Mobile/internal/repositories"
	"fmt"
	"strconv"
	"strings"
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

func (s *Service) GetUsersByParameter(parameters *model.Filter) ([]*model.User, error) {
	var users []*model.User
	var err error
	parametersSql := BuildParametersSql(parameters)
	users, err = s.repo.GetUsersByParameter(parametersSql, parameters.Page)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func BuildParametersSql(parameter *model.Filter) string {
	var columns, value []string

	if parameter.Name != "" {
		columns = append(columns, "name")
		value = append(value, fmt.Sprintf("'%s'", parameter.Name))
	}

	if parameter.SurName != "" {
		columns = append(columns, "surname")
		value = append(value, fmt.Sprintf("'%s'", parameter.SurName))
	}

	if parameter.Patronymic != "" {
		columns = append(columns, "patronymic")
		value = append(value, fmt.Sprintf("'%s'", parameter.Patronymic))
	}

	if parameter.Age != 0 {
		columns = append(columns, "age")
		value = append(value, fmt.Sprintf("%d", parameter.Age))
	}

	if parameter.Gender != "" {
		columns = append(columns, "gender")
		value = append(value, fmt.Sprintf("'%s'", parameter.Gender))
	}

	if parameter.Country != "" {
		columns = append(columns, "country")
		value = append(value, fmt.Sprintf("'%s'", parameter.Country))
	}
	if len(columns) == 0 {
		return ""
	}
	return " where (" + strings.Join(columns, ", ") + ") = (" + strings.Join(value, ", ") + ")"
}
