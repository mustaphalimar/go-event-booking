package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func GenerateToken(userId int64, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		// checking if the token is signed using SigningMethodHS256
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpectde signing method!")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Could not parse token.")
	}

	isTokenValid := parsedToken.Valid

	if !isTokenValid {
		return 0, errors.New("Invalid token.")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("Invalid token claims.")
	}

	userId := int64(claims["userId"].(float64))
	// email := claims["email"].(string)

	return userId, nil
}
