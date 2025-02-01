package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Alistair1231/ipbin/initializers"
	"github.com/gofiber/fiber/v3"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDatabase()
}

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})

	var port = fmt.Sprintf(":%s", os.Getenv("PORT"))
	// Start the server on port 3000
	log.Fatal(app.Listen(port))
}
