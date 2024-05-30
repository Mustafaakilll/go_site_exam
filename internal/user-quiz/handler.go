package userQuiz

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

func (u *UserQuizHandler) GetUserQuizByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	userQuiz, err := u.service.GetUserQuizByID(id)
	if err != nil {
		return err
	}
	return c.JSON(userQuiz)
}

func (u *UserQuizHandler) CreateUserQuiz(c *fiber.Ctx) error {
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

func (u *UserQuizHandler) UpdateUserQuiz(c *fiber.Ctx) error {
	p := new(UpdateUserQuizRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	userQuiz, err := u.service.UpdateUserQuiz(p)
	if err != nil {
		return err
	}
	return c.JSON(userQuiz)
}

func (u *UserQuizHandler) DeleteUserQuiz(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	err = u.service.DeleteUserQuiz(id)
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (u *UserQuizHandler) GetUsersQuizByUserID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("userID")
	if err != nil {
		return err
	}
	p := new(BaseRequest)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	userQuiz, err := u.service.GetUsersQuizzesByUserID(p, id)
	if err != nil {
		return err
	}
	return c.JSON(userQuiz)
}

func (u *UserQuizHandler) GetUsersQuizzesByLessonID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("lessonID")
	if err != nil {
		return err
	}
	p := new(BaseRequest)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	userQuiz, err := u.service.GetUsersQuizzesByLessonID(p, id)
	if err != nil {
		return err
	}
	return c.JSON(userQuiz)
}

func (u *UserQuizHandler) GetUserQuizWithAnswersByUserAndQuizID(c *fiber.Ctx) error {
	quizID, err := c.ParamsInt("quizID")
	if err != nil {
		return err
	}
	userID, err := c.ParamsInt("userID")
	if err != nil {
		return err
	}

	userQuiz, err := u.service.GetUserQuizWithAnswersByUserAndQuizID(userID, quizID)
	if err != nil {
		return err
	}
	return c.JSON(userQuiz)
}

func (u *UserQuizHandler) GetUserQuizByUserAndQuizID(c *fiber.Ctx) error {
	quizID, err := c.ParamsInt("quizID")
	if err != nil {
		return err
	}

	userID, err := c.ParamsInt("userID")
	if err != nil {
		return err
	}
	userQuiz, err := u.service.GetUserQuizByUserAndQuizID(userID, quizID)
	if err != nil {
		return err
	}
	return c.JSON(userQuiz)
}
