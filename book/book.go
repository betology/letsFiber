package book

import (
	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx) error {
	c.SendString("All Books")
}

func GetBook(c *fiber.Ctx) error {
	c.SendString("A Single Book")
}

func NewBook(c *fiber.Ctx) error{
	c.SendString("Adds a new Book")
}
func DeleteBook(c *fiber.Ctx) error{
	 c.SendString("Deletes a Book")
}
