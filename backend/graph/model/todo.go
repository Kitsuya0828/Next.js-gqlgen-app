package model

// import "gorm.io/gorm"

type Todo struct {
	// gorm.Model
	ID     string `gorm:"id" json:"id"`
	Content   string `gorm:"content" json:"content"`
	Done   bool   `gorm:"done" json:"done"`
	UserID string `gorm:"userId" json:"userId"`
	User   *User  `json:"user"`
}
