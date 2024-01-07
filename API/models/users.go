package models

import "github.com/go-playground/validator/v10"

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Age      string `json:"age" validate:"required"`
	HP       string `json:"hp" validate:"required"`
	Email    string `gorm:"unique" validate:"required,email"`
	Password []byte `json:"password" validate:"required,min=6"`
	Ktp      string `json:"ktp" validate:"required"`
	Foto     string `json:"foto" validate:"required"`
}

func (u *User) ValidateUser() error {
	validate := validator.New()
	return validate.Struct(u)
}
