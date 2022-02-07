package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CreateJWTToken create a JWT token
func CreateJWTToken(email string, device string) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	// Set some claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = email
	claims["device"] = device
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
