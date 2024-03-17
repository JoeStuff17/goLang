package productGroup

import (
	"github.com/gofiber/fiber/v2"

	sql "main.go/models"
	productGroup "main.go/services"
)

func CreateGroup(c *fiber.Ctx) error {
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
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Product-Group created Successfully", "data": data.Data})
	} else {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "data": data.Data})
	}
}

func FetchGroup(c *fiber.Ctx) error {
	data := productGroup.FetchAllGroups()
	if data.Success {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Fetched all groups Successfully", "data": data.Data})
	} else {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "data": data.Data})
	}
}

// func UpdateGroup(c *fiber.Ctx) error {
// 	var payload sql.ProductGroup
// 	if err := c.Params(&payload.ID); err != nil {
// 		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
// 			"success": false,
// 			"message": "Not able to process the request",
// 			"data":    err,
// 		})
// 	}
// 	data := productGroup.UpdateById(&payload)
// 	if data.Success {
// 		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Product-Group updated Successfully", "data": data.Data})
// 	} else {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "data": data.Data})
// 	}
// }
