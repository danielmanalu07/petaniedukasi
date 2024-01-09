package models

type Comment struct {
	ID     uint   `json:"id"`
	Text   string `gorm:"type:text" json:"text"`
	UserID uint   `gorm:"" json:"user_id"`
	User   User   `gorm:"foreignKey:UserID" json:"user"`
	PostID uint   `gorm:"" json:"post_id"`
	Post   Post   `gorm:"foreignKey:PostID" json:"post"`
}
