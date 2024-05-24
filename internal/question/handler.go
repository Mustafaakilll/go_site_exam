package question

import (
	"github.com/gofiber/fiber/v2"
)

type QuestionHandler struct {
	service QuestionService
}

func NewQuestionHandler(service *QuestionService) *QuestionHandler {
	return &QuestionHandler{service: *service}
}

func (q *QuestionHandler) GetQuestions(c *fiber.Ctx) error {
	p := new(BaseRequest)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	questions, err := q.service.GetQuestions(p)
	if err != nil {
		return err
	}
	return c.JSON(questions)
}

func (q *QuestionHandler) CreateQuestions(c *fiber.Ctx) error {
	p := new(CreateQuestionRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	questions, err := q.service.CreateQuestions(p)
	if err != nil {
		return err
	}
	return c.JSON(questions)
}
