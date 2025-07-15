package main

import (
	"gostock/backend/core"
	"gostock/backend/core/api"
	"gostock/backend/core/util"
	"log"
	"text/template"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/utils/v2"
)

func main() {
	engine := html.New("./core/template", ".html")
	funcMap := template.FuncMap{
		"formatNumber": util.FormatNumber,
	}
	engine.AddFuncMap(funcMap)
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Define a route
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.Render("index", make(map[string]string))
		if err != nil {
			log.Fatal(err.Error())
			return c.Status(500).SendString("Server Error.")
		}
		return nil
	})

	app.Get("/gostock/:ticker", func(c *fiber.Ctx) error {
		log.Print("Let's analyze stock")
		stockTicker := utils.CopyString(c.Params("ticker"))
		dataCollector := core.NewDataCollector(api.NewAlphaVantageApiProvider())
		data := dataCollector.RequestData(stockTicker)
		log.Print(data)

		err := c.Render("stock", data)
		if err != nil {
			log.Fatal(err.Error())
			return c.Status(500).SendString("Server Error.")
		}
		return nil
	})

	log.Fatal(app.Listen(":3000"))
}
