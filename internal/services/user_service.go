package services

import (
	"Effective_Mobile/internal/model"
	"Effective_Mobile/internal/parsers"
	"Effective_Mobile/internal/repositories"
	"errors"
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

func (s *Service) CreateUser(user *model.User, cfg *model.Config) (*model.User, error) {
	var err error

	user.Age, err = parsers.GetUserAge(user.Name, cfg)
	if err != nil {
		return nil, err
	}

	user.Gender, err = parsers.GetUserGender(user.Name, cfg)
	if err != nil {
		return nil, err
	}

	user.Country, err = parsers.GetUserCountry(user.Name, cfg)
	if err != nil {
		return nil, err
	}

	err = s.repo.SaveUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) UpdateUser(updateUser *model.User) (int64, error) {
	updateInfo, err := BuildUpdateSql(updateUser)
	if err != nil {
		return 0, err
	}
	status, err := s.repo.UpdateUserById(updateUser.Id, updateInfo)
	if err != nil {
		return 0, err
	}

	return status, nil
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

func BuildUpdateSql(updateUser *model.User) (string, error) {
	var updateInfo string
	if updateUser.Name != "" {
		updateInfo += fmt.Sprintf("name = '%s', ", updateUser.Name)
	}

	if updateUser.SurName != "" {
		updateInfo += fmt.Sprintf("surname = '%s', ", updateUser.SurName)
	}

	if updateUser.Patronymic != "" {
		updateInfo += fmt.Sprintf("patronymic = '%s', ", updateUser.Patronymic)
	}

	if updateUser.Age != 0 {
		updateInfo += fmt.Sprintf("age = %d, ", updateUser.Age)
	}

	if updateUser.Gender != "" {
		updateInfo += fmt.Sprintf("gender = '%s', ", updateUser.Gender)
	}

	if updateUser.Country != "" {
		updateInfo += fmt.Sprintf("country = '%s', ", updateUser.Country)
	}
	if updateInfo == "" {
		return "", errors.New("All fields are empty")
	}

	return updateInfo[:len(updateInfo)-2], nil
}
