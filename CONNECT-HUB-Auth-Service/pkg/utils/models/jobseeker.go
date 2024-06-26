package models

import "time"

type JobSeeker struct {
	ID            uint      `json:"id" gorm:"uniquekey; not null"`
	Email         string    `json:"email" gorm:"validate:required"`
	Password      string    `json:"password" gorm:"validate:required"`
	First_name    string    `json:"first_name" gorm:"validate:required"`
	Last_name     string    `json:"last_name" gorm:"validate:required"`
	Phone_number  string    `json:"phone_number" gorm:"validate:required"`
	Date_of_birth string    `json:"date_of_birth" gorm:"validate:required"`
	Gender        string    `json:"gender" gorm:"validate:required"`
	IsBlocked     bool      `json:"is_blocked" gorm:"default:false"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
	Deleted_at    time.Time `json:"deleted_at"`
}
