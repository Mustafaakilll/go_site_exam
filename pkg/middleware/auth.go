package middleware

import (
	"fmt"
	"strings"

	"src/github.com/mustafaakilll/go-site-exam/db"
	"src/github.com/mustafaakilll/go-site-exam/db/entity"
	"src/github.com/mustafaakilll/go-site-exam/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")

	if !strings.HasPrefix(authorization, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Invalid Token"})
	}

	tokenString := strings.TrimPrefix(authorization, "Bearer ")

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		// TODO: move this to config
		return []byte("MyVeryVerySecretKey"), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("invalidate token: %v", err)})
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "invalid token claim"})

	}

	var user entity.User
	db.DB.First(&user, "id = ?", fmt.Sprint(claims["sub"]))

	if utils.IntToString(user.ID) != claims["sub"] {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "the user belonging to this token no logger exists"})
	}

	c.Locals("user", &user)

	return c.Next()
}
