package routes

import (
	"github.com/gofiber/fiber/v2"
	admin_c "main.go/controllers/admin.c"
	"main.go/enums"
	"main.go/middleware"
)

func AdminRoutes(routers fiber.Router) {
	r := routers.Group("admin", middleware.JWTMiddleware())
	r.Post("/create", admin_c.CreateAdmin, middleware.RequireRoles(enums.RoleSuperAdmin))
	r.Get("/fetch-all", admin_c.FetchAllAdmins, middleware.RequireRoles(enums.RoleSuperAdmin, enums.RoleAdmin))
	r.Get("/by-id/:id", admin_c.FetchAdminById, middleware.RequireRoles(enums.RoleSuperAdmin, enums.RoleAdmin))
}
