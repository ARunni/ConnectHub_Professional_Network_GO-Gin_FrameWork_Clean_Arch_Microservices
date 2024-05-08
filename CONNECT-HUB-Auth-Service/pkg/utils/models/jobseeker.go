package models

type JobSeeker struct {
	ID            uint   `json:"id" gorm:"uniquekey; not null"`
	Email         string `json:"email" gorm:"validate:required"`
	Password      string `json:"password" gorm:"validate:required"`
	First_name    string `json:"first_name" gorm:"validate:required"`
	Last_name     string `json:"last_name" gorm:"validate:required"`
	Phone_number  string `json:"phone_number" gorm:"validate:required"`
	Date_of_birth string `json:"date_of_birth" gorm:"validate:required"`
	Gender        string `json:"gender" gorm:"validate:required"`
	Created_at    string `json:"created_at"`
	Updated_at    string `json:"updated_at"`
	Deleted_at    string `json:"deleted_at"`
}
