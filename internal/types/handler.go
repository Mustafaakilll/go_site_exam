package types

import (
	"github.com/gofiber/fiber/v2"
)

type UserTypeHandler struct {
	service UserTypeService
}

func NewUserTypeHandler(service *UserTypeService) *UserTypeHandler {
	return &UserTypeHandler{service: *service}
}

func (u *UserTypeHandler) GetUserTypes(c *fiber.Ctx) error {
	p := new(BaseRequest)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	userTypes, err := u.service.GetUserTypes(p)
	if err != nil {
		return err
	}
	return c.JSON(userTypes)
}

func (u *UserTypeHandler) CreateUserType(c *fiber.Ctx) error {
	p := new(CreateUserTypeRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	userTypes, err := u.service.CreateUserType(p)
	if err != nil {
		return err
	}
	return c.JSON(userTypes)
}
