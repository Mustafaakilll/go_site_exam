package answer

import (
	"github.com/gofiber/fiber/v2"
)

type CodeSubmissionHandler struct {
	service CodeSubmissionService
}

func NewCodeSubmissionHandler(service *CodeSubmissionService) *CodeSubmissionHandler {
	return &CodeSubmissionHandler{service: *service}
}

func (cs *CodeSubmissionHandler) GetCodeSubmissions(c *fiber.Ctx) error {
	p := new(BaseRequest)
	if err := c.QueryParser(p); err != nil {
		return err
	}
	codeSubmissions, err := cs.service.GetCodeSubmissions(p)
	if err != nil {
		return err
	}
	return c.JSON(codeSubmissions)
}

func (cs *CodeSubmissionHandler) CreateCodeSubmissions(c *fiber.Ctx) error {
	p := new(CreateCodeSubmissionRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	codeSubmissions, err := cs.service.CreateCodeSubmissions(p)
	if err != nil {
		return err
	}
	return c.JSON(codeSubmissions)
}
