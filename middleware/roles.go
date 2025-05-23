package middleware

import (
	"github.com/gofiber/fiber/v2"
	"main.go/enums"
	"main.go/models"
)

func RequireRoles(allowedRoles ...enums.UserRole) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*models.Users)

		for _, allowed := range allowedRoles {
			if user.Role == allowed {
				return c.Next()
			}
		}

		// for _, role := range user.Role {
		// 	for _, allowed := range allowedRoles {
		// 		if role == allowed {
		// 			return c.Next()
		// 		}
		// 	}
		// }

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Access denied – insufficient role",
		})
	}
}
