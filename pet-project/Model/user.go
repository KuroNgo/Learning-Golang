package Model

import "time"

type User struct {
	Id       int64     `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Hoten    string    `json:"name"`
	NgaySinh time.Time `json:"ngaysinh"`
	Email    string    `json:"email"`
	SDT      string    `json:"sdt"`
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
