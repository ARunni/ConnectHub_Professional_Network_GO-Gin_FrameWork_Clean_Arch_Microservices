package models

type RecruiterLogin struct {
	Email    string `json:"email" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"min=6,max=20"`
}

type RecruiterSignUp struct {
	Company_name         string `json:"company_name"`
	Industry             string `json:"industry"`
	Company_size         int    `json:"company_size"`
	Website              string `json:"website"`
	Headquarters_address string `json:"headquarters_address"`
	About_company        string `json:"about_company"`
	Contact_email        string `json:"contact_email"`
	Contact_phone_number uint   `json:"contact_phone_number"`
	Password             string `json:"password"`
	ConfirmPassword      string `json:"confirm_password"`
}

type RecruiterDetailsResponse struct {
	ID                   uint   `json:"id"`
	Company_name         string `json:"company_name"`
	Industry             string `json:"industry"`
	Company_size         int    `json:"company_size"`
	Website              string `json:"website"`
	Headquarters_address string `json:"headquarters_address"`
	About_company        string `json:"about_company"`
	Contact_email        string `json:"contact_email"`
	Contact_phone_number uint   `json:"contact_phone_number"`
}

type TokenRecruiter struct {
	Recruiter RecruiterDetailsResponse
	Token     string
}

type RecruiterProfile struct {
	ID                   uint   `json:"id"`
	Company_name         string `json:"company_name"`
	Industry             string `json:"industry"`
	Company_size         int    `json:"company_size"`
	Website              string `json:"website"`
	Headquarters_address string `json:"headquarters_address"`
	About_company        string `json:"about_company"`
	Contact_email        string `json:"contact_email"`
	Contact_phone_number uint   `json:"contact_phone_number"`
}
