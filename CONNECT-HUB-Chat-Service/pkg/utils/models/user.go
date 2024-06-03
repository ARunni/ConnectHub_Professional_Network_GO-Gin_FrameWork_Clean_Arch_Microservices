package models

type UserData struct {
	UserId   uint   `json:"user_id" gorm:"column:id"`
	Username string `json:"username"`
	Profile  string `json:"profile" gorm:"column:imageurl"`
}