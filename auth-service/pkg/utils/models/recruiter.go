package models

type Recruiter struct {
	ID                   uint   `json:"id" gorm:"uniquekey; not null"`
	Company_name         string `json:"company_name" gorm:"validate:required"`
	Industry             string `json:"industry" gorm:"validate:required"`
	Company_size         int    `json:"company_size" gorm:"validate:required"`
	Website              string `json:"website"`
	Headquarters_address string `json:"headquarters_address"`
	About_company        string `json:"about_company" gorm:"type:text"`
	Contact_email        string `json:"contact_email" gorm:"validate:required"`
	Contact_phone_number uint   `json:"contact_phone_number" gorm:"type:numeric"`
	Password             string `json:"password" gorm:"validate:required"`
}