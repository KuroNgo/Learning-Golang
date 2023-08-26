package repo

import "server/model"

func (r *Repo) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}
