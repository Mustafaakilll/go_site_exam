package answer

import "github.com/mustafaakilll/go-site-exam/db/entity"

type BaseRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type BaseResponse struct {
	BaseRequest
	Count int `json:"count"`
}

type CreateCodeSubmissionRequest struct {
	Input          string `json:"input"`
	ExpectedOutput string `json:"expected_output"`

	CodeID int `json:"code_id"`
}

type UpdateCodeSubmissionRequest struct {
	ID             int    `json:"id"`
	Input          string `json:"input"`
	ExpectedOutput string `json:"expected_output"`
}

type CodeSubmissionResponseDTO struct {
	Data []CodeSubmissionDTO `json:"data"`
	BaseResponse
}

type CodeSubmissionDTO struct {
	ID             int         `json:"id"`
	Input          string      `json:"input"`
	ExpectedOutput string      `json:"expected_output"`
	Code           entity.Code `json:"code"`
}
