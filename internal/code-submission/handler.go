package codeSubmission

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

func (cs *CodeSubmissionHandler) CreateCodeSubmission(c *fiber.Ctx) error {
	p := new(CreateCodeSubmissionRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	codeSubmissions, err := cs.service.CreateCodeSubmission(p)
	if err != nil {
		return err
	}
	return c.JSON(codeSubmissions)
}

func (cs *CodeSubmissionHandler) UpdateCodeSubmission(c *fiber.Ctx) error {
	p := new(UpdateCodeSubmissionRequest)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	codeSubmissions, err := cs.service.UpdateCodeSubmission(p)
	if err != nil {
		return err
	}
	return c.JSON(codeSubmissions)
}

func (cs *CodeSubmissionHandler) DeleteCodeSubmission(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	err = cs.service.DeleteCodeSubmission(id)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"message": "Code Submission deleted successfully"})
}
