package middleware

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"main.go/enums"
	dto "main.go/interface_model"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")
		if authorization == "" || len(authorization) < 8 || authorization[:7] != "Bearer " {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "Insufficient access: Missing or invalid Authorization header",
			})
		}

		tokenStr := authorization[7:]

		token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Optional: check token.Method here if you want to enforce "HS256" only
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_KEY")), nil // Use env vars in production
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "Unauthorized: Invalid token claims",
			})
		}

		expirationFloat, ok := claims["exp"].(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "Unauthorized: Invalid token expiration",
			})
		}

		idStr, ok := claims["id"].(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "Unauthorized: Invalid token ID",
			})
		}

		idInt, _ := strconv.Atoi(idStr)
		role, _ := claims["role"].(string)
		name, _ := claims["name"].(string)

		c.Locals("user", dto.ReqUser{
			ID:   uint(idInt),
			Role: enums.UserRole(role),
			Name: name,
		})

		expiration := time.Unix(int64(expirationFloat), 0)
		if time.Now().After(expiration) {
			log.Println("Token has expired")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "Unauthorized: Token has expired",
			})
		}

		return c.Next()
	}
}
