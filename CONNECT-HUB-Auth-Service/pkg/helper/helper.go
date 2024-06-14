package helper

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"regexp"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to generate password hash")
	}
	return string(hashPassword), nil
}

func CompareHashAndPassword(a string, b string) error {
	err := bcrypt.CompareHashAndPassword([]byte(a), []byte(b))
	if err != nil {
		return err
	}
	return nil
}

func ValidatePhoneNumber(phone string) bool {
	phoneNumber := phone
	pattern := `^\d{10}$`
	regex := regexp.MustCompile(pattern)
	value := regex.MatchString(phoneNumber)
	return value
}

func GenerateVideoCallKey(userID, oppositeUser int) (string, error) {
	currentTime := strconv.FormatInt(time.Now().UnixNano(), 10)
	key := strconv.Itoa(userID) + "_" + strconv.Itoa(oppositeUser) + "_" + currentTime
	hash := md5.Sum([]byte(key))
	keyString := hex.EncodeToString(hash[:])

	return keyString, nil
}
