package helper

import (
	"github.com/ARunni/connectHub_gateway/pkg/config"
	"fmt"

	"github.com/golang-jwt/jwt"
)

type authCustomClaimsAdmin struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	jwt.StandardClaims
}

func ValidateTokenAdmin(tokenString string) (*authCustomClaimsAdmin, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error at validate token config loading")
		return nil, err
	}
	// Define a function to retrieve the HMAC key.
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		// Ensure the token is signed with the HMAC algorithm.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the HMAC key.
		return []byte(cfg.AdminAccessKey), nil
	}

	// Parse the token with custom claims.
	token, err := jwt.ParseWithClaims(tokenString, &authCustomClaimsAdmin{}, keyFunc)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	// Validate the token and extract custom claims.
	if claims, ok := token.Claims.(*authCustomClaimsAdmin); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
