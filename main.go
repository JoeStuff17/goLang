package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"main.go/database"
	"main.go/routes"
)

func main() {
	port := os.Getenv("PORT")
	// version := os.Getenv("VERSION")

	if port == "" {
		port = "3000"
	}
	// if version == "" {
	// 	version = "v1"
	// }

	database.ConnectToMySql()
	app := fiber.New()
	app.Use(cors.New())
	app.Use(helmet.New())

	// api := app.Group("/api/" + version)
	api := app.Group("/api/")
	routes.ProductGroupRoutes(api)
	log.Fatal(app.Listen(":" + port))
}
