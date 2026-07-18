package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"landing-page/views"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return component.Render(c.Context(), c.Response().BodyWriter())
}

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using default environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := fiber.New(fiber.Config{
		// Custom global error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			
			if code == fiber.StatusNotFound {
				c.Status(fiber.StatusNotFound)
				lang := "vi"
				if strings.HasPrefix(c.Path(), "/en") {
					lang = "en"
				}
				return render(c, views.NotFound(lang))
			}

			c.Status(code)
			return c.SendString(fmt.Sprintf("Internal Server Error: %v", err))
		},
	})

	// Middlewares
	app.Use(logger.New(logger.Config{
		Format: `{"ip":"${ip}", "time":"${time}", "status":${status}, "latency":"${latency}", "method":"${method}", "path":"${path}"}` + "\n",
	}))
	app.Use(compress.New())
	app.Use(cors.New())

	// Static files routing
	app.Static("/static", "./public")

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return render(c, views.Index("vi"))
	})

	app.Get("/en", func(c *fiber.Ctx) error {
		return render(c, views.Index("en"))
	})

	// Fallback 404 handler for routes that are not registered
	app.Use(func(c *fiber.Ctx) error {
		return fiber.ErrNotFound
	})

	// Start server
	log.Printf("Server starting on port %s in %s mode...\n", port, os.Getenv("ENV_MODE"))
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
