package userAnswer

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
	userAnswers, err := u.service.CreateUserAnswer(p)
	if err != nil {
		return err
	}
	return c.JSON(userAnswers)
}

func (u *UserAnswerHandler) UpdateUserAnswer(c *fiber.Ctx) error {
	p := new(UpdateUserAnswerRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	userAnswers, err := u.service.UpdateUserAnswer(p)
	if err != nil {
		return err
	}
	return c.JSON(userAnswers)
}

func (u *UserAnswerHandler) DeleteUserAnswer(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	err = u.service.DeleteUserAnswer(id)
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (u *UserAnswerHandler) GetUserAnswerByQuestionID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("questionID")
	if err != nil {
		return err
	}
	userAnswer, err := u.service.GetUserAnswerByQuestionID(id)
	if err != nil {
		return err
	}
	return c.JSON(userAnswer)
}

func (u *UserAnswerHandler) GetUserAnswerByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	userAnswer, err := u.service.GetUserAnswerByID(id)
	if err != nil {
		return err
	}
	return c.JSON(userAnswer)
}
