package answer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mustafaakilll/go-site-exam/pkg/utils"
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

func (a *AnswerHandler) GetAnswerByID(c *fiber.Ctx) error {
	id := c.Params("id")
	answer, err := a.service.GetAnswerByID(utils.StringToInt(id))
	if err != nil {
		return err
	}
	return c.JSON(answer)
}

func (a *AnswerHandler) GetAnswersByQuestionID(c *fiber.Ctx) error {
	p := new(BaseRequest)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	questionID, err := c.ParamsInt("question_id")
	if err != nil {
		return err
	}
	answers, err := a.service.GetAnswersByQuestionID(p, questionID)
	if err != nil {
		return err
	}
	return c.JSON(answers)
}
