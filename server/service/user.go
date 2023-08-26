package service

import "server/repo"

type User struct {
	repo repo.IRepo
}

func (s *User) NewUser(repo repo.IRepo) *User {
	return &User{
		repo: repo,
	}
}

func (s *User) CreateUser() error {
	return nil
}

func (s *User) GetUser() error {
	return nil
}

func (s *User) GetUserByID() error {
	return nil
}

func (s *User) UpdateUser() error {
	return nil
}
