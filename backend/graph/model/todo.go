package model

type Todo struct {
	ID     string `json:"id"`
	Content   string `json:"content"`
	Done   bool   `json:"done"`
	UserID string `json:"userId"`
	User   *User  `json:"user"`
}
