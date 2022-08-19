package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome!!!")
}

func main() {
	app := fiber.New()

	app.Get("/api", welcome)

	log.Fatal(app.Listen(":3000"))
}