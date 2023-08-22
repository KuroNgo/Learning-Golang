package service

import (
	"context"
	"errors"

	"erp-server/model"
	"erp-server/repo"
)

type User struct {
	repo repo.IRepo
}

type IUser interface {
	Login(ctx context.Context, userRequest model.UserRequest) (userResponse model.User, err error)
	Register(ctx context.Context, userRequest model.User) (userResponse model.User, err error)
}

func NewUser(repo repo.IRepo) *User {
	return &User{
		repo: repo,
	}
}

func (s *User) Login(ctx context.Context, userRequest model.UserRequest) (userResponse model.User, err error) {
	user, err := s.repo.GetUserByEmail(userRequest.Username)
	if err != nil {
		return model.User{}, err
	}
	// kiểm tra xem mật khẩu có đúng không
	if user.Password != userRequest.Password {
		return model.User{}, errors.New("Invalid username or password")
	}
	return user, nil
}

// Done but ID not going up
func (s *User) Register(ctx context.Context, userRegister model.User) (userResponse model.User, err error) {
	baseModel := model.BaseModel{}
	userRequest := model.User{
		BaseModel:   baseModel,
		Username:    userRegister.Username,
		Password:    userRegister.Password,
		FullName:    userRegister.FullName,
		Address:     userRegister.Address,
		DateOfBirth: userRegister.DateOfBirth,
		Email:       userRegister.Email,
		Role:        userRegister.Role,
		Phone:       userRegister.Phone,
	}
	user, err := s.repo.CreateUser(userRequest)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
