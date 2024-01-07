package models

import "github.com/go-playground/validator/v10"

type Edukasi struct {
	Id          uint   `json:"id"`
	Title       string `json:"title" validate:"required" gorm:"not null"`
	Image       string `json:"image" validate:"required" gorm:"not null"`
	Description string `json:"description" validate:"required" gorm:"not null"`
}

func (u *Edukasi) ValidateEdukasi() error {
	validate := validator.New()
	return validate.Struct(u)
}
