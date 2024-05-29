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
	questions, err := q.service.CreateQuestion(p)
	if err != nil {
		return err
	}
	return c.JSON(questions)
}

func (q *QuestionHandler) UpdateQuestion(c *fiber.Ctx) error {
	p := new(UpdateQuestionRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	questions, err := q.service.UpdateQuestion(p)
	if err != nil {
		return err
	}
	return c.JSON(questions)
}

func (q *QuestionHandler) DeleteQuestion(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	err = q.service.DeleteQuestion(id)
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (q *QuestionHandler) GetQuestionsByQuizID(c *fiber.Ctx) error {
	p := new(BaseRequest)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	quizID, err := c.ParamsInt("quiz_id")
	if err != nil {
		return err
	}
	questions, err := q.service.GetQuestionsByQuizID(p, quizID)
	if err != nil {
		return err
	}
	return c.JSON(questions)
}

func (q *QuestionHandler) GetQuestionByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	question, err := q.service.GetQuestionByID(id)
	if err != nil {
		return err
	}
	return c.JSON(question)
}

// func (q *QuestionHandler) GetQuestionsWithChoices(c *fiber.Ctx) error {
// 	p := new(BaseRequest)
// 	if err := c.QueryParser(p); err != nil {
// 		return err
// 	}
// 	questionID, err := c.ParamsInt("quizID")
// 	if err != nil {
// 		return err
// 	}
// 	questions, err := q.service.GetQuestionsWithChoices(questionID)
// 	if err != nil {
// 		return err
// 	}
// 	return c.JSON(questions)
// }
