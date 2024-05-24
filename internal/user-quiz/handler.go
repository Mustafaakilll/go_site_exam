package user_quiz

import (
	"github.com/gofiber/fiber/v2"
)

type UserQuizHandler struct {
	service UserQuizService
}

func NewUserQuizHandler(service *UserQuizService) *UserQuizHandler {
	return &UserQuizHandler{service: *service}
}

func (u *UserQuizHandler) GetUserQuizzes(c *fiber.Ctx) error {
	p := new(BaseRequest)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	userQuizzes, err := u.service.GetUserQuizzes(p)
	if err != nil {
		return err
	}
	return c.JSON(userQuizzes)
}

func (u *UserQuizHandler) CreateUserQuizzes(c *fiber.Ctx) error {
	p := new(CreateUserQuizRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	userQuiz, err := u.service.CreateUserQuizzes(p)
	if err != nil {
		return err
	}
	return c.JSON(userQuiz)
}
