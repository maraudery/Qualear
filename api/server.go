package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/maraudery/ra/api/handlers"
	"github.com/maraudery/ra/api/route"
)

// Launch starts the API
func main() {
	app := fiber.New(fiber.Config{
		// Pass view engine
		Views: html.New("./views", ".html"),
		// Pass global error handler
		ErrorHandler: handlers.Errors("./public/500.html"),
	})

	// Render index template with IP input value
	app.Get("/", handlers.Render())

	// Serve static assets
	app.Static("/", "./public", fiber.Static{
		Compress:      true,
		CacheDuration: -1,
		MaxAge:        0, // this is redundant as 0 is the default
	})

	// Main GEO handler that is cached for 10 minutes
	route.Username(app)
	route.Ip(app)
	route.Discord(app)
	route.Vin(app)

	// Handle 404 errors
	app.Use(handlers.NotFound("./public/404.html"))

	// Listen on port :6174
	log.Fatal(app.Listen(":6174"))
}
