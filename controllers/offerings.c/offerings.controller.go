package offerings_c

import (
	"github.com/gofiber/fiber/v2"

	dto "main.go/interface_model"
	"main.go/models"
	offerings_s "main.go/services/offering.s"
)

func CreateOffering(c *fiber.Ctx) error {
	payload := new(models.Offerings)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Not able to process the request 2",
			"data":    err.Error(),
		})
	}
	localUser := c.Locals("user").(dto.ReqUser)
	res := offerings_s.CreateOffering(payload, localUser)
	return c.Status(res.StatusCode).JSON(fiber.Map{"success": res.Success, "message": res.Message, "data": res.Data})
}

func FetchChurchOfferings(c *fiber.Ctx) error {
	payload := new(dto.OfferingsFetchPayload)
	if err := c.QueryParser(payload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Not able to process the request 2",
			"data":    err.Error(),
		})
	}
	println("church_id", *payload)
	res := offerings_s.FetchChurchOfferings(payload.ChurchId)
	return c.Status(res.StatusCode).JSON(fiber.Map{"success": res.Success, "message": res.Message, "data": res.Data, "count": res.Count})
}

func FetchOfferingsByMember(c *fiber.Ctx) error {
	payload := new(dto.FetchOfferingsByMemberPayload)
	if err := c.QueryParser(payload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Not able to process the request 2",
			"data":    err.Error(),
		})
	}
	res := offerings_s.FetchOfferingsByMember(int(payload.ChurchId), int(payload.MemberId))
	return c.Status(res.StatusCode).JSON(fiber.Map{"success": res.Success, "message": res.Message, "data": res.Data})
}
