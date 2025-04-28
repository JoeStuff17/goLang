package routes

import (
	"github.com/gofiber/fiber/v2"
	church_c "main.go/controllers/church.c"
	"main.go/enums"
	"main.go/middleware"
)

func ChurchRoutes(routers fiber.Router) {
	r := routers.Group("church", middleware.JWTMiddleware())
	r.Post("/create", church_c.CreateChurch, middleware.RequireRoles(enums.RoleSuperAdmin, enums.RoleAdmin))
}
