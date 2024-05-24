package quiz

import (
	"github.com/gofiber/fiber/v2"
)

type QuizHandler struct {
	service QuizService
}

func NewQuizHandler(service *QuizService) *QuizHandler {
	return &QuizHandler{service: *service}
}

func (q *QuizHandler) GetQuizzes(c *fiber.Ctx) error {
	p := new(BaseRequest)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	quizzes, err := q.service.GetQuizzes(p)
	if err != nil {
		return err
	}
	return c.JSON(quizzes)
}

func (q *QuizHandler) CreateQuizzes(c *fiber.Ctx) error {
	p := new(CreateQuizRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	quizzes, err := q.service.CreateQuizzes(p)
	if err != nil {
		return err
	}
	return c.JSON(quizzes)
}
