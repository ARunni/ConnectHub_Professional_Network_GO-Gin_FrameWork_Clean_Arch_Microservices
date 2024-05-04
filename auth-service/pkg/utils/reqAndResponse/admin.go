package reqandresponse

type AdminLogin struct {
	Email    string `json:"email" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"required"`
}

type AdminDetailsResponse struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
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
