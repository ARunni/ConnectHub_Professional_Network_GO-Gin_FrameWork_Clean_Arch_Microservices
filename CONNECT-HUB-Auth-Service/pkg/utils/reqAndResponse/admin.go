package reqandresponse

import "github.com/ARunni/ConnetHub_auth/pkg/utils/models"

type AdminLogin struct {
	Email    string `json:"email" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"required"`
}

type AdminDetailsResponse struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
type AdminSignUp struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type TokenAdmin struct {
	Admin AdminDetailsResponse
	Token string
}

type AdminLoginRes struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JobseekerDetailsAtAdmin struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email" validate:"email"`
	Phone   string `json:"phone"`
	Blocked bool   `json:"blocked"`
}
type RecruiterDetailsAtAdmin struct {
	Id           int    `json:"id"`
	CompanyName  string `json:"company_name"`
	Contact_mail string `json:"contact_mail" validate:"email"`
	Phone        string `json:"phone"`
	Blocked      bool   `json:"blocked"`
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
	Policies models.Policy `json:"polices"`
}

type GetAllPolicyRes struct {
	Policies []models.Policy `json:"polices"`
}
