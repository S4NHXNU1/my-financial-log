package controller

import "github.com/gofiber/fiber/v2"

func PageController(app fiber.Router) {

	app.Get("", func(c *fiber.Ctx) error {
		return c.Status(200).SendFile("./web/static/main.html")
	})

	app.Get("/usd", func(c *fiber.Ctx) error {
		return c.Status(200).SendFile("./web/static/usd.html")
	})

	app.Get("/gold", func(c *fiber.Ctx) error {
		return c.Status(200).SendFile("./web/static/gold.html")
	})

}
