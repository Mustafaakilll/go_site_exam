package user_answer

import "src/github.com/mustafaakilll/go-site-exam/db/entity"

type BaseRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type BaseResponse struct {
	BaseRequest
	Count int `json:"count"`
}

type CreateUserAnswerRequest struct {
	UserID   int `json:"user_id"`
	QuizID   int `json:"quiz_id"`
	AnswerID int `json:"answer_id"`
}

type UserAnswerResponseDTO struct {
	Data []UserAnswerDTO `json:"data"`
	BaseResponse
}

type UserAnswerDTO struct {
	ID     int           `json:"id"`
	User   entity.User   `json:"user"`
	Quiz   entity.Quiz   `json:"quiz"`
	Answer entity.Answer `json:"answer"`
}
