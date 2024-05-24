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

type CreateAnswerRequest struct {
	Text string `json:"text"`

	QuestionID int `json:"question_id"`
	ChoiceID   int `json:"choice_id"`
	UserID     int `json:"user_id"`
}

type UpdateAnswerRequest struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type AnswerResponseDTO struct {
	Data []AnswerDTO `json:"data"`
	BaseResponse
}

type AnswerDTO struct {
	ID       int             `json:"id"`
	Text     string          `json:"text"`
	Question entity.Question `json:"question"`
	Choice   entity.Choice   `json:"choice"`
	User     entity.User     `json:"user"`
}
