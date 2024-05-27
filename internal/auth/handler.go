package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mustafaakilll/go-site-exam/internal/user"
)

type AuthHandler struct {
	service AuthService
}

func NewAuthHandler(service *AuthService) *AuthHandler {
	return &AuthHandler{service: *service}
}

func (u *AuthHandler) Login(c *fiber.Ctx) error {
	p := new(LoginRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	token, err := u.service.Login(p)
	if err != nil {
		return err
	}
	return c.JSON(token)
}

func (u *AuthHandler) Register(c *fiber.Ctx) error {
	p := new(user.CreateUserRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	token, user, err := u.service.RegisterUser(p)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"token": token,
		"user":  user,
	})
}
