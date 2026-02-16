package auth

import (
	"strings"
	"gostock/backend/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"

)

func AuthMiddleware(c *fiber.Ctx) error {
	tokenStr := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})
	if err != nil || !token.Valid {
		logger.Log.Error("Failed to validate token",
			zap.String("token", tokenStr),
			zap.Error(err),
		)
		return c.SendStatus(401)
	}
	return c.Next()
}
