package models

import "time"

type Admin struct {
	ID        uint   `json:"id" gorm:"uniquekey; not null"`
	Firstname string `json:"firstname" gorm:"validate:required"`
	Lastname  string `json:"lastname" gorm:"validate:required"`
	Email     string `json:"email" gorm:"validate:required"`
	Password  string `json:"password" gorm:"validate:required"`
}
type AdminLogin struct {
	Email    string `json:"email" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"min=6,max=20"`
}
type AdminSignUp struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type AdminDetailsResponse struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"Email"`
}
type TokenAdmin struct {
	Admin AdminDetailsResponse
	Token string
}

type JobseekerDetailsAtAdmin struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email" validate:"email"`
	Phone   string `json:"phone"`
	Blocked bool   `json:"blocked"`
}
type RecruiterDetailsAtAdmin struct {
	Id          int    `json:"id"`
	CompanyName string `json:"company_name"`
	Email       string `json:"contact_mail" validate:"email"`
	Phone       string `json:"phone"`
	Blocked     bool   `json:"blocked"`
}

type BlockRes struct {
	Status string `json:"status"`
}

type CreatePolicyReq struct {
	Title   string `json:"title" gorm:"validate:required"`
	Content string `json:"content" gorm:"validate:required"`
}
type UpdatePolicyReq struct {
	Id      int    `json:"id"`
	Title   string `json:"title" gorm:"validate:required"`
	Content string `json:"content" gorm:"validate:required"`
}

type CreatePolicyRes struct {
	Policies Policy `json:"polices"`
}

type GetAllPolicyRes struct {
	Policies []Policy `json:"polices"`
}

type Policy struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Title     string    `json:"title" gorm:"validate:required"`
	Content   string    `json:"content" gorm:"validate:required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
