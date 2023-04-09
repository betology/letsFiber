package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/betology/letsfiber/book"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func setupRoute(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)

func main() {
    app := fiber.New()

    app.Get("/", helloWorld)

    app.Listen(":3000")
}
