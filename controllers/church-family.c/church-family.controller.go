package church_family_c

import (
	"github.com/gofiber/fiber/v2"

	dto "main.go/interface_model"
	church_family_s "main.go/services/church-family.s"
)

func FetchChurchFamilies(c *fiber.Ctx) error {
	payload := new(dto.ChurchFamiliesFetchPayload)
	if err := c.QueryParser(payload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Not able to process the request 2",
			"data":    err.Error(),
		})
	}
	res := church_family_s.GetAllChurchFamilies(int(payload.ChurchId))
	return c.Status(res.StatusCode).JSON(fiber.Map{"success": res.Success, "message": res.Message, "data": res.Data, "count": res.Count})
}

func FetchChurchFamilyById(c *fiber.Ctx) error {
	payload := new(dto.ChurchFamilyFetchPayload)
	if err := c.QueryParser(payload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Not able to process the request 2",
			"data":    err.Error(),
		})
	}
	res := church_family_s.FetchChurchFamilyById(int(payload.ChurchId), int(payload.FamilyId))
	return c.Status(res.StatusCode).JSON(fiber.Map{"success": res.Success, "message": res.Message, "data": res.Data})
}
