package models

type Like struct {
	ID     uint `json:"id"`
	UserID uint `gorm:"" json:"user_id"`
	User   User `gorm:"foreignKey:UserID" json:"user"`
	PostID uint `gorm:"" json:"post_id"`
	Post   Post `gorm:"foreignKey:PostID" json:"post"`
	Status int  `json:"status"`
}
