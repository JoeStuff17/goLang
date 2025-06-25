package routes

import (
	"github.com/gofiber/fiber/v2"
	user_c "main.go/controllers/user.c"
)

func UserRoutes(routers fiber.Router) {
	r := routers.Group("user")
	r.Post("/create", user_c.CreateUser)
	r.Post("/send-login-otp", user_c.SendOtp)
	r.Post("/verify-login-otp", user_c.VerifyOtp)
	r.Post("/send-mail", user_c.MailSendTest)
}
