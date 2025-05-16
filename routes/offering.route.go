package routes

import (
	"github.com/gofiber/fiber/v2"
	offerings_c "main.go/controllers/offerings.c"
	"main.go/enums"
	"main.go/middleware"
)

func OfferingRoutes(routers fiber.Router) {
	r := routers.Group("offerings", middleware.JWTMiddleware())
	r.Post("/create", middleware.RequireRoles(enums.RoleSuperAdmin, enums.RoleAdmin), offerings_c.CreateOffering)
	r.Get("/fetch", middleware.RequireRoles(enums.RoleSuperAdmin, enums.RoleAdmin), offerings_c.FetchChurchOfferings)
	r.Get("/get-by-member", middleware.RequireRoles(enums.RoleSuperAdmin, enums.RoleAdmin), offerings_c.FetchOfferingsByMember)
}
