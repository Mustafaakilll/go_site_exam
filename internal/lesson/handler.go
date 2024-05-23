package lesson

import (
	"github.com/gofiber/fiber/v2"
)

type LessonHandler struct {
	service LessonService
}

func NewLessonHandler(service *LessonService) *LessonHandler {
	return &LessonHandler{service: *service}
}

func (l *LessonHandler) GetLessons(c *fiber.Ctx) error {
	p := new(BaseRequest)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	lessons, err := l.service.GetLessons(p)
	if err != nil {
		return err
	}
	return c.JSON(lessons)
}

func (l *LessonHandler) CreateLessons(c *fiber.Ctx) error {
	p := new(CreateLessonRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	lessons, err := l.service.CreateLessons(p)
	if err != nil {
		return err
	}
	return c.JSON(lessons)
}
