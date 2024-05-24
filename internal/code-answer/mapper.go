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

type CreateCodeAnswerRequest struct {
	Status string `json:"status"`

	UserID int `json:"user_id"`
	CodeID int `json:"code_id"`
}

type UpdateCodeAnswerRequest struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

type CodeAnswerResponseDTO struct {
	Data []CodeAnswerDTO `json:"data"`
	BaseResponse
}

type CodeAnswerDTO struct {
	ID     int         `json:"id"`
	Status string      `json:"status"`
	User   entity.User `json:"user"`
	Code   entity.Code `json:"code"`
}
