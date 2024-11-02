package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type jwtClaims struct {
	Email            string `json:"email"`
	RegisteredClaims jwt.RegisteredClaims
}

func CreateAccessToken(email, secret string) (string, error) {
	key := []byte(secret)

	jwtClaims := jwtClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // TODO: вынести в конфиг
			Subject:   email,
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims.RegisteredClaims)
	return accessToken.SignedString(key)
}

func VerifyAccessToken(accessTokenString string, secret string) (*jwt.Token, error) {
	accessToken, err := jwt.Parse(accessTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func AuthenticateUser(email, hashedPassword, password, secret string) (string, error) {
	if err := VerifyPassword(hashedPassword, password); err != nil {
		return "", err
	}

	accessToken, err := CreateAccessToken(email, secret)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
