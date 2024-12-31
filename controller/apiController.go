package controller

import (
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v2"
)

type USDs struct {
	THB float32 `json:"thb"`
	USD float32 `json:"usd"`
}

type Golds struct {
	USD float32 `json:"usd"`
	BG  float32 `json:"bg"`
}

type Earnings struct {
	THB float32 `json:"thb"`
}

func ApiController(app fiber.Router) {
	const USDfilePath = "./configs/data-usd.json"
	const GoldfilePath = "./configs/data-gold.json"
	const EarningsfilePath = "./configs/data-earnings.json"

	app.Get("/usd", func(c *fiber.Ctx) error {
		_usd := []USDs{{THB: 0.0, USD: 0.0}}

		jsonData, err := os.ReadFile(USDfilePath)
		if err != nil {
			_jsonData, err := json.MarshalIndent(_usd, "", "  ")
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "Failed to marshal default count"})
			}

			err = os.WriteFile(USDfilePath, _jsonData, 0644)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "Failed to create count.json"})
			}

			return c.Status(201).JSON(_usd)
		}

		err = json.Unmarshal(jsonData, &_usd)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to parse count.json"})
		}

		return c.Status(200).JSON(_usd)
	})

	app.Post("/usd", func(c *fiber.Ctx) error {
		var _usd []USDs

		if err := c.BodyParser(&_usd); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Failed to parse JSON",
			})
		}

		_jsonData, err := json.MarshalIndent(_usd, "", "  ")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to marshal default count"})
		}

		err = os.WriteFile(USDfilePath, _jsonData, 0644)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create count.json"})
		}

		return c.Status(200).JSON(_usd)
	})

	app.Get("/gold", func(c *fiber.Ctx) error {
		_gold := []Golds{{USD: 0.0, BG: 0.0}}

		jsonData, err := os.ReadFile(GoldfilePath)
		if err != nil {
			_jsonData, err := json.MarshalIndent(_gold, "", "  ")
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "Failed to marshal default count"})
			}

			err = os.WriteFile(GoldfilePath, _jsonData, 0644)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "Failed to create count.json"})
			}

			return c.Status(201).JSON(_gold)
		}

		err = json.Unmarshal(jsonData, &_gold)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to parse count.json"})
		}

		return c.Status(200).JSON(_gold)
	})

	app.Post("/gold", func(c *fiber.Ctx) error {
		var _gold []Golds

		if err := c.BodyParser(&_gold); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Failed to parse JSON",
			})
		}

		_jsonData, err := json.MarshalIndent(_gold, "", "  ")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to marshal default count"})
		}

		err = os.WriteFile(GoldfilePath, _jsonData, 0644)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create count.json"})
		}

		return c.Status(200).JSON(_gold)
	})

	app.Get("/earning", func(c *fiber.Ctx) error {
		_earn := []Earnings{{THB: 0.0}}

		jsonData, err := os.ReadFile(EarningsfilePath)
		if err != nil {
			_jsonData, err := json.MarshalIndent(_earn, "", "  ")
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "Failed to marshal default count"})
			}

			err = os.WriteFile(EarningsfilePath, _jsonData, 0644)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "Failed to create count.json"})
			}

			return c.Status(201).JSON(_earn)
		}

		err = json.Unmarshal(jsonData, &_earn)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to parse count.json"})
		}

		return c.Status(200).JSON(_earn)
	})

	app.Post("/earning", func(c *fiber.Ctx) error {
		var _earn []Earnings

		if err := c.BodyParser(&_earn); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Failed to parse JSON",
			})
		}

		_jsonData, err := json.MarshalIndent(_earn, "", "  ")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to marshal default count"})
		}

		err = os.WriteFile(EarningsfilePath, _jsonData, 0644)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create count.json"})
		}

		return c.Status(200).JSON(_earn)
	})
}
