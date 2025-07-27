package main

import (
	"gostock/backend/core"
	"gostock/backend/core/api"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/utils/v2"
)

func main() {

	app := fiber.New()

	// cors config
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://127.0.0.1:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Global or route-specific limiter
	app.Use(limiter.New(limiter.Config{
		Max:        20,        // Allow max 20 requests
		Expiration: time.Hour, // Per an hour
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // Rate limit by IP
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Rate limit exceeded. Try again later.",
			})
		},
	}))

	app.Get("/api/:symbol", func(c *fiber.Ctx) error {
		stockTicker := utils.CopyString(c.Params("symbol"))
		dataCollector := core.NewDataCollector(api.NewAlphaVantageApiProvider())
		data := dataCollector.RequestData(stockTicker)
		return c.JSON(data)
	})

	log.Fatal(app.Listen(":8080"))
}
