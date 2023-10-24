package repositories

import (
	"Effective_Mobile/internal/model"
	"github.com/jmoiron/sqlx"
)

type Repo interface {
	UserRepo
}

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repo {
	return &Repository{
		db: db,
	}
}

type UserRepo interface {
	SaveUser(messageInfo *model.User) error
	GetInfoAllUsers(page string) ([]*model.User, error)
	DeleteUserById(id string) error
	UpdateUserById(updateUser *model.User) error
}
