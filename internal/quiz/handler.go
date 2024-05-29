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

	userID := c.Locals("user_id").(int)
	quizzes, err := q.service.CreateQuizzes(p, userID)
	if err != nil {
		return err
	}
	return c.JSON(quizzes)
}

func (q *QuizHandler) UpdateQuiz(c *fiber.Ctx) error {
	p := new(UpdateQuizRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	quizzes, err := q.service.UpdateQuiz(p)
	if err != nil {
		return err
	}
	return c.JSON(quizzes)
}

func (q *QuizHandler) DeleteQuiz(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	err = q.service.DeleteQuiz(id)
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (q *QuizHandler) GetQuizByID(c *fiber.Ctx) error {
	p := new(BaseRequest)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	quiz, err := q.service.GetQuizByID(id)
	if err != nil {
		return err
	}
	return c.JSON(quiz)
}

func (q *QuizHandler) GetQuizByTeacher(c *fiber.Ctx) error {
	p := new(BaseRequest)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	teacherID, err := c.ParamsInt("teacher_id")
	if err != nil {
		return err
	}
	quizzes, err := q.service.GetQuizByTeacher(p, teacherID)
	if err != nil {
		return err
	}
	return c.JSON(quizzes)
}

func (q *QuizHandler) GetJoinedUserByQuizID(c *fiber.Ctx) error {
	quizID, err := c.ParamsInt("quizID")
	if err != nil {
		return err
	}
	quizzes, err := q.service.GetJoinedUserByQuizID(quizID)
	if err != nil {
		return err
	}
	return c.JSON(quizzes)
}
