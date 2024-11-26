package main

import "github.com/gofiber/fiber/v2" // add Fiber package

func main() {
	app := fiber.New() // create a new Fiber instance

	// Create a new endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("Hello, World!") // send text
	})

	// Start server on port 3000
	app.Listen(":8080")
}
