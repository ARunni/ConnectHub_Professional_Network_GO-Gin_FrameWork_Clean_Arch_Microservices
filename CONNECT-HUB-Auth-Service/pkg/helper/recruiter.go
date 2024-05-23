package helper

import (
	"ConnetHub_auth/pkg/config"
	req "ConnetHub_auth/pkg/utils/reqAndResponse"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type authCustomClaimsRecruiter struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func GenerateTokenRecruiter(recruiter req.RecruiterDetailsResponse) (string, error) {
	claims := &authCustomClaimsRecruiter{
		Id:    recruiter.ID,
		Email: recruiter.Contact_email,
		Role:  "recruiter",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.RecruiterAccessKey))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}
	return tokenString, nil
}
