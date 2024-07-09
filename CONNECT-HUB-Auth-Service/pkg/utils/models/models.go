package models

import "time"

type UserDatas struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
}

type Users struct {
	ID                   uint      `json:"id" gorm:"uniquekey; not null"`
	Email                string    `json:"email" gorm:"validate:required"`
	Password             string    `json:"password" gorm:"validate:required"`
	Role                 string    `json:"role" gorm:"role:2;default:'Jobseeker';Check:role IN ('Jobseeker', 'Recruiter');"`
	First_name           string    `json:"first_name" gorm:"validate:required"`
	Last_name            string    `json:"last_name" gorm:"validate:required"`
	Phone_number         string    `json:"phone_number" gorm:"validate:required"`
	Gender               string    `json:"gender" gorm:"validate:required"`
	Date_of_birth        string    `json:"date_of_birth" gorm:"validate:required"`
	IsBlocked            bool      `json:"is_blocked" gorm:"default:false"`
	Industry             string    `json:"industry" gorm:"validate:required"`
	Company_size         int       `json:"company_size" gorm:"validate:required"`
	Website              string    `json:"website" gorm:"default:NA"`
	Headquarters_address string    `json:"headquarters_address"`
	About_company        string    `json:"about_company" gorm:"type:text"`
	Created_at           time.Time `json:"created_at"`
	Updated_at           time.Time `json:"updated_at"`
	Deleted_at           time.Time `json:"deleted_at"`
}
