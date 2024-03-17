package routes

import (
	productGroup "main.go/controllers"

	"github.com/gofiber/fiber/v2"
)

func ProductGroupRoutes(routers fiber.Router) {
	r := routers.Group("product-group")
	r.Post("/", productGroup.CreateGroup)
	r.Get("/", productGroup.FetchGroup)
	// r.Put("/", productGroup.UpdateGroup)
}
