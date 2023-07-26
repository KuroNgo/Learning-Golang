package entity

type Person struct {
	FirstName string `json:"first-name" binding:"required"`
	LastName  string `json:"last-name" bindind:"required"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" validate:"required,email"`
}
