package main

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/wlopezob/gomiprimerapirest01/api"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			if err, ok := e.(*fiber.Error); ok {
				return errors.New(fmt.Sprintf("this is a error fiber %s",
					err.Message))
			}
			return errors.New("this is a managed error")
		},
	})
	//midleware
	app.Use(recover.New())
	app.Use(logger.New())

	api.RoutesUp(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 03👋!")
	})

	app.Listen(":3000")
}
