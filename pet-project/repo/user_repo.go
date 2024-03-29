package repo

import (
	"pet-project/model"
)

type IUserRepo interface {
	GetUserByEmail(email string) (model.User, error)
}

func (repo *Repo) GetUserByEmail(email string) (model.User, error) {
	// Tìm kiếm bản ghi trong bảng "users" từ request
	var user model.User
	repo.db.Where("email = ?", email).First(&user)
	return user, nil
}
