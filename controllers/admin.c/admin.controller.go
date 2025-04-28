package admin_c

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	dto "main.go/interface_model"
	"main.go/models"
	admin_s "main.go/services/admin.s"
)

func CreateAdmin(c *fiber.Ctx) error {
	var payload models.Admins
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Not able to process the request",
			"data":    err,
		})
	}
	localUser := c.Locals("user").(dto.ReqUser)
	res := admin_s.CreateAdmin(payload, localUser)
	return c.Status(res.StatusCode).JSON(fiber.Map{"success": res.Success, "message": res.Message, "data": res.Data})
}

func FetchAllAdmins(c *fiber.Ctx) error {
	res := admin_s.FetchAllAdmins()
	return c.Status(res.StatusCode).JSON(fiber.Map{"success": res.Success, "message": res.Message, "data": res.Data})

}

func FetchAdminById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	res := admin_s.FetchAdminById(id)
	return c.Status(res.StatusCode).JSON(fiber.Map{"success": res.Success, "message": res.Message, "data": res.Data})

}
