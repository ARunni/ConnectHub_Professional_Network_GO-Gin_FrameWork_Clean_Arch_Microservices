package helper

import (
	"github.com/ARunni/ConnetHub_auth/pkg/config"
	req "github.com/ARunni/ConnetHub_auth/pkg/utils/reqAndResponse"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type authCustomClaimsJobseeker struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func GenerateTokenJobseeker(jobseeker req.JobSeekerDetailsResponse) (string, error) {
	claims := &authCustomClaimsJobseeker{
		Id:    jobseeker.ID,
		Email: jobseeker.Email,
		Role:  "jobseeker",
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
	tokenString, err := token.SignedString([]byte(cfg.JobSeekerAccessKey))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}
	return tokenString, nil
}
