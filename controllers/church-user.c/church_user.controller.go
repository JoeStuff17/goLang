package church_user_c

import (
	"github.com/gofiber/fiber/v2"

	dto "main.go/interface_model"
	"main.go/models"
	church_user_s "main.go/services/church-user.s"
)

func CreateChurchUser(c *fiber.Ctx) error {
	payload := new(models.ChurchUser)
	if err := c.BodyParser(payload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Not able to process the request - controller",
			"data":    err.Error(),
		})
	}
	localUser := c.Locals("user").(dto.ReqUser)
	res := church_user_s.CreateChurchUser(payload, localUser)
	return c.Status(res.StatusCode).JSON(fiber.Map{"success": res.Success, "message": res.Message, "data": res.Data})
}

func FetchChurchUsers(c *fiber.Ctx) error {
	payload := new(dto.ChurchUsersFetchPayload)
	if err := c.QueryParser(payload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Not able to process the request 2",
			"data":    err.Error(),
		})
	}
	res := church_user_s.GetAllChurchUsers(int(payload.ChurchId))
	return c.Status(res.StatusCode).JSON(fiber.Map{"success": res.Success, "message": res.Message, "data": res.Data, "count": res.Count})
}

func FetchChurchUserById(c *fiber.Ctx) error {
	payload := new(dto.ChurchUserFetchPayload)
	if err := c.QueryParser(payload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"success": false,
			"message": "Not able to process the request 2",
			"data":    err.Error(),
		})
	}
	res := church_user_s.FetchChurchUserById(payload.ChurchId, payload.UserId)
	return c.Status(res.StatusCode).JSON(fiber.Map{"success": res.Success, "message": res.Message, "data": res.Data})
}
