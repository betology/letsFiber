package book

import (
	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx) error {
	return c.Send("All Books")
}

func GetBook(c *fiber.Ctx) error {
	return c.Send("A Single Book")
}

func NewBook(c *fiber.Ctx) error{
	return c.Send("Adds a new Book")
}
func DeleteBook(c *fiber.Ctx) error{
	return c.Send("Deletes a Book")
}
