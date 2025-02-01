package main

import (
	"log"
	"os"

	"github.com/Alistair1231/ipbin/controllers"
	"github.com/Alistair1231/ipbin/initializers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDatabase()
	initializers.SyncDB()
}

func main() {
	engine := html.New("./views", ".html")
	// Setup App
	app := fiber.New(
		fiber.Config{
			Views: engine,
		},
	)

	// Configure App
	app.Static("/", "./public")

	// Routes
	app.Get("/", controllers.GetIndex)

	var port = ":" + os.Getenv("PORT")
	// Start app
	log.Fatal(app.Listen(port))
}
