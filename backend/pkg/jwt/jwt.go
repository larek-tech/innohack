package jwt

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/larek-tech/innohack/backend/internal/auth/model"
	"github.com/larek-tech/innohack/backend/internal/shared"
	"golang.org/x/crypto/bcrypt"
)

type jwtClaims struct {
	Email            string `json:"email"`
	RegisteredClaims jwt.RegisteredClaims
}

func CreateAccessToken(userID int64, email, secret string) (string, error) {
	key := []byte(secret)

	jwtClaims := jwtClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // TODO: вынести в конфиг
			Subject:   strconv.FormatInt(userID, 10),
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

func AuthenticateUser(userID int64, email, hashedPassword, password, secret string) (model.TokenResp, error) {
	if err := VerifyPassword(hashedPassword, password); err != nil {
		return model.TokenResp{}, shared.ErrInvalidCredentials
	}

	accessToken, err := CreateAccessToken(userID, email, secret)
	if err != nil {
		return model.TokenResp{}, err
	}

	return model.TokenResp{
		Token: accessToken,
		Type:  shared.BearerType,
	}, nil
}
