package models

import "github.com/go-playground/validator/v10"

type Admin struct {
	Id       uint   `json:"id"`
	Name     string `json:"name" validate:"required"`
	Email    string `gorm:"unique" validate:"required,email"`
	Password []byte `json:"password" validate:"required"`
}

func (u *Admin) ValidateAdmin() error {
	validate := validator.New()
	return validate.Struct(u)
}
