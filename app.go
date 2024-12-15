package main

import (
	"my-financial-log/configs"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.AppSetup(app)
	app.Listen(":3000")
}
