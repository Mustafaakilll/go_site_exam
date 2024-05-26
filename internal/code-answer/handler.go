package codeAnswer

import (
	"github.com/gofiber/fiber/v2"
)

type CodeAnswerHandler struct {
	service CodeAnswerService
}

func NewCodeAnswerHandler(service *CodeAnswerService) *CodeAnswerHandler {
	return &CodeAnswerHandler{service: *service}
}

func (ca *CodeAnswerHandler) GetCodeAnswers(c *fiber.Ctx) error {
	p := new(BaseRequest)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	codeAnswers, err := ca.service.GetCodeAnswers(p)
	if err != nil {
		return err
	}
	return c.JSON(codeAnswers)
}

func (ca *CodeAnswerHandler) CreateCodeAnswer(c *fiber.Ctx) error {
	p := new(CreateCodeAnswerRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	codeAnswers, err := ca.service.CreateCodeAnswers(p)
	if err != nil {
		return err
	}
	return c.JSON(codeAnswers)
}

func (ca *CodeAnswerHandler) UpdateCodeAnswer(c *fiber.Ctx) error {
	p := new(UpdateCodeAnswerRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	codeAnswers, err := ca.service.UpdateCodeAnswer(p)
	if err != nil {
		return err
	}
	return c.JSON(codeAnswers)
}

func (ca *CodeAnswerHandler) DeleteCodeAnswer(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	err = ca.service.DeleteCodeAnswer(id)
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}
