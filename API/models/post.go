package models

import "github.com/go-playground/validator/v10"

type Post struct {
	ID      uint   `json:"id"`
	Content string `gorm:"type:varchar(225)" json:"content" validate:"required"`
	Image   string `gorm:"type:varchar(225)" json:"image" validate:"required"`
	UserID  uint   `gorm:"" json:"user_id"`
	User    User   `gorm:"foreignKey:UserID" json:"user"`
}

func (p *Post) ValidatePost() error {
	validate := validator.New()
	return validate.Struct(p)
}
