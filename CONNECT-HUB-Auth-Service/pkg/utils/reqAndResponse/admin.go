package reqandresponse

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
	Id          int    `json:"id"`
	CompanyName string `json:"company_name"`
	Email       string `json:"contact_mail" validate:"email"`
	Phone       string `json:"phone"`
	Blocked     bool   `json:"blocked"`
}

type BlockRes struct {
	Status string `json:"status"`
}
