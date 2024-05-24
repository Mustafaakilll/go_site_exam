package choice

import (
	"github.com/gofiber/fiber/v2"
)

type ChoiceHandler struct {
	service ChoiceService
}

func NewChoiceHandler(service *ChoiceService) *ChoiceHandler {
	return &ChoiceHandler{service: *service}
}

func (c *ChoiceHandler) GetChoices(ctx *fiber.Ctx) error {
	p := new(BaseRequest)
	if err := ctx.QueryParser(p); err != nil {
		return err
	}
	choices, err := c.service.GetChoices(p)
	if err != nil {
		return err
	}
	return ctx.JSON(choices)
}

func (c *ChoiceHandler) CreateChoices(ctx *fiber.Ctx) error {
	p := new(CreateChoiceRequest)
	if err := ctx.BodyParser(p); err != nil {
		return err
	}
	choices, err := c.service.CreateChoices(p)
	if err != nil {
		return err
	}
	return ctx.JSON(choices)
}
