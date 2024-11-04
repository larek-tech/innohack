package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/larek-tech/innohack/backend/internal/shared"
	"github.com/larek-tech/innohack/backend/pkg/jwt"
	"go.opentelemetry.io/otel/trace"
)

const (
	unauthorized = "unauthorized"
)

func Jwt(secret string, tracer trace.Tracer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, span := tracer.Start(c.Context(), "auth.middleware.jwt")
		defer span.End()

		authHeader := c.Get("Authorization", unauthorized)
		if authHeader == unauthorized {
			return shared.ErrMissingJwt
		}

		t := strings.Split(authHeader, " ")
		if len(t) != 2 {
			return shared.ErrMissingJwt
		}

		authToken := t[1]
		token, err := jwt.VerifyAccessToken(authToken, secret)
		if err != nil {
			return shared.ErrInvalidJwt
		}

		userID, err := token.Claims.GetSubject()
		if err != nil {
			return shared.ErrInvalidJwt
		}

		c.Locals(shared.UserIDKey, userID)

		return c.Next()
	}
}
