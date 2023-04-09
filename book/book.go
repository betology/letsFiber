package book

import (
	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx) {
	c.SendString("All Books")
}

func GetBook(c *fiber.Ctx) {
	c.SendString("A Single Book")
}

func NewBook(c *fiber.Ctx) {
	c.SendString("Adds a new Book")
}
func DeleteBook(c *fiber.Ctx) {
	c.SendString("Deletes a Book")
}
