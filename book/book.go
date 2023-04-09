package book

import (
	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx) {
	s.SendString("All Books")
}

func GetBook(c *fiber.Ctx) {
	s.SendString("A Single Book")
}

func NewBook(c *fiber.Ctx) {
	s.SendString("Adds a new Book")
}
func DeleteBook(c *fiber.Ctx) {
	s.SendString("Deletes a Book")
}
