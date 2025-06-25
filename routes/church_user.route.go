package routes

import (
	"github.com/gofiber/fiber/v2"
	church_user_c "main.go/controllers/church-user.c"
	"main.go/enums"
	"main.go/middleware"
)

func ChurchUserRoutes(routers fiber.Router) {
	r := routers.Group("church-user", middleware.JWTMiddleware())
	r.Post("/create", middleware.RequireRoles(enums.RoleSuperAdmin, enums.RoleAdmin), church_user_c.CreateChurchUser)
	r.Get("/fetch-all", middleware.RequireRoles(enums.RoleSuperAdmin, enums.RoleAdmin), church_user_c.FetchChurchUsers)
	r.Get("/by-id", middleware.RequireRoles(enums.RoleSuperAdmin, enums.RoleAdmin), church_user_c.FetchChurchUserById)
}
