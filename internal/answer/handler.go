package answer

import (
	"github.com/gofiber/fiber/v2"
)

type AnswerHandler struct {
	service AnswerService
}

func NewAnswerHandler(service *AnswerService) *AnswerHandler {
	return &AnswerHandler{service: *service}
}

func (a *AnswerHandler) GetAnswers(c *fiber.Ctx) error {
	p := new(BaseRequest)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	answers, err := a.service.GetAnswers(p)
	if err != nil {
		return err
	}
	return c.JSON(answers)
}

func (a *AnswerHandler) CreateAnswer(c *fiber.Ctx) error {
	p := new(CreateAnswerRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	answers, err := a.service.CreateAnswers(p)
	if err != nil {
		return err
	}
	return c.JSON(answers)
}

func (a *AnswerHandler) UpdateAnswer(c *fiber.Ctx) error {
	p := new(UpdateAnswerRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	answers, err := a.service.UpdateAnswers(p)
	if err != nil {
		return err
	}
	return c.JSON(answers)
}

func (a *AnswerHandler) DeleteAnswer(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	err = a.service.DeleteAnswer(id)
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}
