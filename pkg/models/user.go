package models

type User struct {
	FirstName string `json:"firstName" validate:"required,lte=10"`
	LastName  string `json:"lastName" validate:"required,lte=10"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,gte=8,lte=64"`
}

type Login struct {
	Email     string `json:"emailIn" validate:"required,email"`
	Password  string `json:"passwordIn" validate:"required,gte=8,lte=64"`
}