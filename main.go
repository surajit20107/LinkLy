package main

import (
	"log"
	"fmt"
	"os"
	"github.com/gofiber/fiber/v2"
	"main/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil{
		fmt.Println("Warning: Error loading .env, using system environment variables")
	}
	
	app := fiber.New()

	routes.UserRoute(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Server is up and running...!! 🚀",
			"author": "Surajit",
			"github": "https://github.com/surajit20107",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}
	
	log.Fatal(app.Listen(port))
}
