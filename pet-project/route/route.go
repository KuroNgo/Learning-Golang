package route

import (
	"pet-project/config"
	repo2 "pet-project/repo"
	"pet-project/service"
)

type Service struct {
	*config.App
}

func NewService() *Service {
	s := Service{
		config.NewApp(),
	}

	db := s.GetDB()
	repo := repo2.NewRepo(db)

	userService := service.NewUser(repo)
	user := handler.NewUser
}
