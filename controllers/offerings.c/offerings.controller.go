package offerings_c

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	dto "main.go/interface_model"
	"main.go/models"
	offerings_s "main.go/services/offering.s"
	user_s "main.go/services/user.s"
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
	res := offerings_s.FetchChurchOfferings(int(payload.ChurchId))
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
	localUser := c.Locals("user").(dto.ReqUser)
	userRes := user_s.FetchUserById(int(localUser.ID))
	fmt.Println(userRes.Data)
	if !userRes.Success || userRes.Data == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": "Unable to fetch user details"})
	}
	userData, ok := userRes.Data.(dto.FetchUserWithDetails)
	fmt.Println(userData)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Unexpected user data format"})
	}
	role := userData.Role
	fmt.Println("User role:", role)
	if role != "church_admin" && role != "church_user" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": "Unauthorized access"})
	}

	res := offerings_s.FetchOfferingsByMember(int(*userData.ChurchId), int(userData.RelatedId))
	return c.Status(res.StatusCode).JSON(fiber.Map{"success": res.Success, "message": res.Message, "data": res.Data})
}
