package user_c

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"main.go/helpers"
	dto "main.go/interface_model"
	"main.go/models"
	user_s "main.go/services/user.s"
)

func CreateUser(c *fiber.Ctx) error {
	var payload models.Users
	fmt.Println("payload------>", payload)
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Not able to process the request",
			"data":    err,
		})
	}
	data := user_s.CreateUser(&payload)
	if data.Success {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "User created Successfully", "data": data.Data})
	} else {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "data": data.Data})
	}
}

func AdminSendOtp(c *fiber.Ctx) error {
	payload := new(dto.AdminSendOtpPayload)
	if err := c.BodyParser(payload); err != nil {
		return err
	}

	res := user_s.AdminSendOtp(payload)
	return c.Status(res.StatusCode).JSON(fiber.Map{"success": res.Success, "data": res.Data, "message": res.Message})
}

func AdminVerifyOtp(c *fiber.Ctx) error {
	payload := new(dto.AdminVerifyOtpPayload)
	if err := c.BodyParser(payload); err != nil {
		return err
	}

	res := user_s.AdminVerifyOtp(payload)
	return c.Status(res.StatusCode).JSON(fiber.Map{"success": res.Success, "data": res.Data, "message": res.Message})
}

func MailSendTest(c *fiber.Ctx) error {
	payload := new(dto.AdminVerifyOtpPayload)
	if err := c.BodyParser(payload); err != nil {
		return err
	}

	helpers.SendOtpMail()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "data": nil, "message": "mail send successfully"})
}
