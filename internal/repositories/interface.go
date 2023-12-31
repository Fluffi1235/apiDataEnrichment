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
	GetUsersByParameter(parameters string, page int) ([]*model.User, error)
	UpdateUserById(id int, updateInfo string) (int64, error)
	DeleteUserById(id int) error
}
