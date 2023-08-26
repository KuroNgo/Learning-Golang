package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;"`
	Email    string    `gorm:"type:varchar(100);unique_index;not null"`
	Password string    `gorm:"type:varchar(100);not null"`
	Name     string    `gorm:"type:varchar(100);not null"`

	Role  string `gorm:"type:varchar(100);not null"`
	Photo string `gorm:"type:varchar(100);not null"`

	Verified  bool      `gorm:"default:false"`
	Provider  string    `gorm:"default:'local'"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (user *User) BeforeCreate(*gorm.DB) error {
	user.ID = uuid.NewV4()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return nil
}

// The RegisterUserInput struct will be used to validate
// the request payload during the signup process,
// while the LoginUserInput struct will be used to validate
// the request payload during the sign-in process.
type RegisterUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Role      string `json:"role,omitempty"`
	Provider  string `json:"provider,omitempty"`
	Photo     string `json:"photo,omitempty"`
	Verified  bool   `json:"verified,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FilteredResponse(user *User) UserResponse {
	return UserResponse{
		ID:        user.ID.String(),
		Email:     user.Email,
		Name:      user.Name,
		Role:      user.Role,
		Provider:  user.Provider,
		Photo:     user.Photo,
		Verified:  user.Verified,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
