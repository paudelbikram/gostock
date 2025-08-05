package main

import (
	"gostock/backend/core"
	"gostock/backend/core/api"
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

	app := fiber.New()

	// Add recovery middleware to catch panics
	// Let's log these
	app.Use(recover.New())

	// cors config
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://127.0.0.1:3000, http://192.168.1.70:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Global or route-specific limiter
	app.Use(limiter.New(limiter.Config{
		Max:        20,               // Allow max 20 requests
		Expiration: 10 * time.Minute, // Per 10 minutes
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // Rate limit by IP
		},
		LimitReached: func(c *fiber.Ctx) error {
			logger.Log.Error("Rate limit exceeded.",
				zap.String("ip", c.IP()),
			)
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Rate limit exceeded. Try again later.",
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

	app.Get("/api/:symbol", func(c *fiber.Ctx) error {
		stockTicker := utils.CopyString(c.Params("symbol"))
		if stockTicker == "" {
			// Return HTTP 400 with error JSON
			logger.Log.Error("No Ticker Symbol provide")
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "stock symbol is required",
			})
		}
		dataCollector := core.NewDataCollector(api.NewAlphaVantageApiProvider())
		data, err := dataCollector.RequestData(strings.ToUpper(stockTicker))
		if err != nil {
			logger.Log.Error("internal server error",
				zap.Error(err),
			)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "internal server error occurred",
			})
		}
		return c.JSON(data)
	})

	logger.Log.Fatal("Backend starting...", zap.Error(app.Listen(":8080")))
}
