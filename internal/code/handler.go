package code

import (
	"github.com/gofiber/fiber/v2"
)

type CodeHandler struct {
	service CodeService
}

func NewCodeHandler(service *CodeService) *CodeHandler {
	return &CodeHandler{service: *service}
}

func (h *CodeHandler) GetCodes(c *fiber.Ctx) error {
	p := new(BaseRequest)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	codes, err := h.service.GetCodes(p)
	if err != nil {
		return err
	}
	return c.JSON(codes)
}

func (h *CodeHandler) CreateCode(c *fiber.Ctx) error {
	p := new(CreateCodeRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	codes, err := h.service.CreateCode(p)
	if err != nil {
		return err
	}
	return c.JSON(codes)
}

func (h *CodeHandler) UpdateCode(c *fiber.Ctx) error {
	p := new(UpdateCodeRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	codes, err := h.service.UpdateCode(p)
	if err != nil {
		return err
	}
	return c.JSON(codes)
}

func (h *CodeHandler) DeleteCode(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	err = h.service.DeleteCode(id)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"message": "Code deleted successfully"})
}

func (h *CodeHandler) GetCodesByLessonID(c *fiber.Ctx) error {
	p := new(BaseRequest)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	id, err := c.ParamsInt("lessonID")
	if err != nil {
		return err
	}
	codes, err := h.service.GetCodesByLessonID(p, id)
	if err != nil {
		return err
	}
	return c.JSON(codes)
}

func (h *CodeHandler) GetCodeByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	code, err := h.service.GetCodeByID(id)
	if err != nil {
		return err
	}
	return c.JSON(code)
}

func (h *CodeHandler) GetCodesByTeacherID(c *fiber.Ctx) error {
	p := new(BaseRequest)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	id, err := c.ParamsInt("teacherID")
	if err != nil {
		return err
	}
	codes, err := h.service.GetCodesByTeacherID(p, id)
	if err != nil {
		return err
	}
	return c.JSON(codes)
}

func (h *CodeHandler) GetUsersCodes(c *fiber.Ctx) error {
	p := new(BaseRequest)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	id, err := c.ParamsInt("userID")
	if err != nil {
		return err
	}
	codes, err := h.service.GetUsersCodes(p, id)
	if err != nil {
		return err
	}
	return c.JSON(codes)
}
