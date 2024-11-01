package handler

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) OAuth(c *fiber.Ctx) error {
	// http://localhost:9999/auth/oauth#access_token=y0_AgAAAAAfdxvpAAy0AwAAAAEWoi6GAADCgqlVfQ9Dvp6WiZQPjZPcyhsuWw&token_type=bearer&expires_in=31536000&cid=v7pk2e57t3rzutztp10q3739kw

	requestURI := string(c.Path())

	parts := strings.Split(requestURI, "#")
	if len(parts) != 2 {
		return errors.New("invalid params")
	}

	values := strings.Split(parts[1], "&")

	requestJSON := make(map[string]string)
	for _, v := range values {
		kv := strings.Split(v, "=")
		if len(kv) != 2 {
			continue
		}
		requestJSON[kv[0]] = kv[1]
	}
	fmt.Println(requestJSON)

	token, err := h.service.OauthGetAccessToken(c.Context(), requestJSON["access_token"])
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"code": token})
}
