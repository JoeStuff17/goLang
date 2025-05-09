package routes

import (
	"github.com/gofiber/fiber/v2"
	church_c "main.go/controllers/church.c"
	"main.go/enums"
	"main.go/middleware"
)

func ChurchFamilyRoutes(routers fiber.Router) {
	r := routers.Group("church-family", middleware.JWTMiddleware())
	r.Get("/fetch-all", middleware.RequireRoles(enums.RoleSuperAdmin, enums.RoleAdmin), church_c.FetchChurches)
}
