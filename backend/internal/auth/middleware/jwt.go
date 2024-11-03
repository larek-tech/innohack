package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/shared"
	"github.com/larek-tech/innohack/backend/pkg"
	"github.com/larek-tech/innohack/backend/pkg/jwt"
)

const (
	unauthorized = "unauthorized"
)

func Jwt(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization", unauthorized)
		if authHeader == unauthorized {
			return pkg.WrapErr(shared.ErrMissingJwt)
		}

		t := strings.Split(authHeader, " ")
		if len(t) != 2 {
			return pkg.WrapErr(shared.ErrMissingJwt)
		}

		authToken := t[1]
		token, err := jwt.VerifyAccessToken(authToken, secret)
		if err != nil {
			return pkg.WrapErr(shared.ErrInvalidJwt, err.Error())
		}

		userID, err := token.Claims.GetSubject()
		if err != nil {
			return pkg.WrapErr(shared.ErrInvalidJwt, err.Error())
		}

		c.Locals(shared.UserIDKey, userID)

		return c.Next()
	}
}
