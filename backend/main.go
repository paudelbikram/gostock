package main

import (
	"context"
	"fmt"
	"gostock/backend/auth"
	"gostock/backend/config"
	"gostock/backend/core"
	"gostock/backend/core/api"
	"gostock/backend/core/util"
	"gostock/backend/logger"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/utils/v2"
	"go.uber.org/zap"
)

func main() {
	// Setting up logger
	logger.Init()
	defer logger.Sync()

	// Loading config
	config, err := config.NewConfig()
	if err != nil {
		return // exit if failed to load config
	}

	// Initializing firebase
	auth.InitFirebase()

	app := fiber.New()

	// Add recovery middleware to catch panics
	// Use recover with custom handler
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		// Custom error handler
		StackTraceHandler: func(c *fiber.Ctx, err interface{}) {
			logger.Log.Error("Panic recovered",
				zap.String("path", c.OriginalURL()),
				zap.Any("Error", err),
			)
			// Send JSON response with 500
			err = c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal server error. Please try again later.",
			})
		},
	}))

	// cors config
	app.Use(cors.New(cors.Config{
		AllowOrigins: config.CORSOrigin,
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowCredentials: true,
	}))

	// Global or route-specific limiter
	app.Use(limiter.New(limiter.Config{
		Max:        15,              // Allow max 15 requests
		Expiration: 1 * time.Minute, // Per 1 minutes
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // Rate limit by IP
		},
		LimitReached: func(c *fiber.Ctx) error {
			logger.Log.Error("Rate limit exceeded.",
				zap.String("ip", c.IP()),
			)
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": "Rate limit exceeded. Please slow down.",
			})
		},
	}))

	// Middleware to log requests
	app.Use(func(c *fiber.Ctx) error {
		logger.Log.Info("API Request",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.String("ip", c.IP()),
		)
		return c.Next()
	})

	app.Post("/auth/login", func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		client, _ := auth.FirebaseApp.Auth(context.Background())
		token, err := client.VerifyIDToken(context.Background(), tokenStr)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "invalid token"})
		}
		email := token.Claims["email"].(string)
		name := token.Claims["name"]
		jwt, _ := auth.CreateJWT(email)
		return c.JSON(fiber.Map{
			"user": fiber.Map{
				"email": email,
				"name":  name,
				"token": jwt,
			},
		})
	})

	app.Post("/api/stock/list", auth.AuthMiddleware, func(c *fiber.Ctx) error {
		list, err := util.GetCacheStock()
		if err != nil {
			return c.JSON([]string{})
		}
		return c.JSON(list)
	})

	app.Post("/api/:symbol", auth.AuthMiddleware, func(c *fiber.Ctx) error {
		stockTicker := utils.CopyString(c.Params("symbol"))
		if stockTicker == "" {
			// Return HTTP 400 with error JSON
			logger.Log.Error("No Ticker Symbol provide")
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Stock symbol is required.",
			})
		}
		dataCollector := core.NewDataCollector(api.NewAlphaVantageApiProvider())
		data, err := dataCollector.RequestData(strings.ToUpper(stockTicker))
		if err != nil {
			logger.Log.Error("internal server error",
				zap.Error(err),
			)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal server error occurred. Please try again later.",
			})
		}
		return c.JSON(data)
	})

	logger.Log.Fatal("Backend starting...", zap.Error(app.Listen(":"+fmt.Sprint(config.Port))))
}
