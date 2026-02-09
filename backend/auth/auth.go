package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *fiber.Ctx) error {
	tokenStr := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})
	if err != nil || !token.Valid {
		return c.SendStatus(401)
	}
	return c.Next()
}
