package helper

import (
	"ConnetHub_auth/pkg/config"
	req "ConnetHub_auth/pkg/utils/reqAndResponse"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type authCustomClaimsAdmin struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateTokenAdmin(admin req.AdminDetailsResponse) (string, error) {
	claims := &authCustomClaimsAdmin{
		Id:    admin.ID,
		Email: admin.Email,
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
	tokenString, err := token.SignedString([]byte(cfg.AdminAccessKey))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}
	return tokenString, nil
}
