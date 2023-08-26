package repo

import (
	"gorm.io/gorm"

	"server/model"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{
		db: db,
	}
}

type IRepo interface {
	CreateUser(user *model.User) error
}
