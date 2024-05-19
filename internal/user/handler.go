package user

import (
	"src/github.com/mustafaakilll/go-site-exam/pkg/models"
	"src/github.com/mustafaakilll/go-site-exam/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: *service}
}

func (u *UserHandler) GetUsers(c *fiber.Ctx) error {
	p := new(models.PaginateRequest)
	if err := c.QueryParser(p); err != nil {
		return err
	}

	users, err := u.service.GetUsers(p)
	if err != nil {
		return err
	}
	return c.JSON(users)
}

func (u *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := u.service.GetUserByID(utils.StringToInt(id))
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (u *UserHandler) CreateUser(c *fiber.Ctx) error {
	p := new(CreateUserRequest)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	err := u.service.CreateUser(p)
	if err != nil {
		return err
	}
	return nil

}

func (u *UserHandler) UpdateUser(c *fiber.Ctx) error {
	p := new(UpdateUserRequest)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	err := u.service.UpdateUser(p)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	return u.service.DeleteUser(utils.StringToInt(id))
}

func (u *UserHandler) GetUserByEmail(c *fiber.Ctx) error {
	email := c.Params("email")

	user, err := u.service.GetUserByEmail(email)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (u *UserHandler) GetUserByUsername(c *fiber.Ctx) error {
	username := c.Params("username")

	user, err := u.service.GetUserByUsername(username)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (u *UserHandler) SetTeacher(c *fiber.Ctx) error {
	id := c.Params("id")

	return u.service.SetTeacher(utils.StringToInt(id))
}
