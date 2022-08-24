package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/home", fiber.Map{
			"Title": "Hello World",
		})
	})

	err := godotenv.Load()
	if err != nil {
		log.Fatal(" error loading .env")
	}

	port := os.Getenv("PORT")
	app.Listen(":" + port)
}
