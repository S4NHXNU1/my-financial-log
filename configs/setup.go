package configs

import (
	"my-financial-log/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func AppSetup(app *fiber.App) {
	app.Use(cors.New())

	apiGroup := app.Group("/api")
	controller.ApiController(apiGroup)

	pageGroup := app.Group("/")
	controller.PageController(pageGroup)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendFile("./web/template/404.html")
	})
}
