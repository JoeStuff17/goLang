package routes

import (
	"github.com/gofiber/fiber/v2"
	user_c "main.go/controllers/user.c"
)

func UserRoutes(routers fiber.Router) {
	r := routers.Group("user")
	r.Post("/create", user_c.CreateUser)
	r.Post("/send-login-otp/admin", user_c.AdminSendOtp)
	r.Post("/verify-login-otp/admin", user_c.AdminVerifyOtp)
	r.Post("/send-mail", user_c.MailSendTest)
}
