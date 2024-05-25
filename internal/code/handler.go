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

func (h *CodeHandler) CreateCodes(c *fiber.Ctx) error {
	p := new(CreateCodeRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	codes, err := h.service.CreateCodes(p)
	if err != nil {
		return err
	}
	return c.JSON(codes)
}
