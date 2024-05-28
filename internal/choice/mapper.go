package choice

import "github.com/mustafaakilll/go-site-exam/db/entity"

type BaseRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type BaseResponse struct {
	BaseRequest
	Count int `json:"count"`
}

type CreateChoiceRequest struct {
	Text      string `json:"text"`
	IsCorrect bool   `json:"is_correct"`

	QuestionID int `json:"question_id"`
}

type UpdateChoiceRequest struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	IsCorrect bool   `json:"is_correct"`
}

type ChoiceResponseDTO struct {
	Data []ChoiceDTO `json:"data"`
	BaseResponse
}

type ChoiceWithQuestionResponseDTO struct {
	BaseResponse
	Data []ChoiceDTO `json:"data"`
}

type ChoiceWithQuestionDTO struct {
	ID        int             `json:"id"`
	Text      string          `json:"text"`
	IsCorrect bool            `json:"is_correct"`
	Question  entity.Question `json:"question"`
}

type ChoiceDTO struct {
	ID        int             `json:"id"`
	Text      string          `json:"text"`
	IsCorrect bool            `json:"is_correct"`
	Question  entity.Question `json:"question"`
}
