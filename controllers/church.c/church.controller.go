package church_c

import (
	"github.com/gofiber/fiber/v2"

	dto "main.go/interface_model"
	"main.go/models"
	church_s "main.go/services/church.s"
)

func CreateChurch(c *fiber.Ctx) error {
	var payload models.Churches
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Not able to process the request",
			"data":    err,
		})
	}
	localUser := c.Locals("user").(dto.ReqUser)
	data := church_s.CreateChurch(payload, localUser)
	if data.Success {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Church created Successfully", "data": data.Data})
	} else {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "data": data.Data})
	}
}
