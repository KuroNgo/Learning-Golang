package entity

type Login struct {
	username string `json:"login" binding:"required"`
	password string `json:"password" binding:"required"`
}
