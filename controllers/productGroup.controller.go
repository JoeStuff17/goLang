package productGroup

import (
	"github.com/gofiber/fiber/v2"

	sql "main.go/models"
)

func CreateProductGroup(c *fiber.Ctx) error {
	var payload sql.ProductGroup
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Not able to process the request",
			"data":    err,
		})
	}
	data := productGroup.CreateProductGroup(&payload)
	if data.Success {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "data": data.Data, "message": "Wrapper Got Executed"})
	} else {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "data": data.Data})
	}
}
