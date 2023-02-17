package models

type User struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
