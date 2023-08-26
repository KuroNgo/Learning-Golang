package service

import "server/repo"

type Service struct {
	repo repo.IRepo
}

type IService interface {
	NewUser(repo repo.IRepo) *User
	CreateUser() error
	GetUser() error
	GetUserByID() error
	UpdateUser() error
}
