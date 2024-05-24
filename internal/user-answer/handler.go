package user_answer

import (
	"github.com/gofiber/fiber/v2"
)

type UserAnswerHandler struct {
	service UserAnswerService
}

func NewUserAnswerHandler(service *UserAnswerService) *UserAnswerHandler {
	return &UserAnswerHandler{service: *service}
}

func (u *UserAnswerHandler) GetUserAnswers(c *fiber.Ctx) error {
	p := new(BaseRequest)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	userAnswers, err := u.service.GetUserAnswers(p)
	if err != nil {
		return err
	}
	return c.JSON(userAnswers)
}

func (u *UserAnswerHandler) CreateUserAnswers(c *fiber.Ctx) error {
	p := new(CreateUserAnswerRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	userAnswers, err := u.service.CreateUserAnswers(p)
	if err != nil {
		return err
	}
	return c.JSON(userAnswers)
}
