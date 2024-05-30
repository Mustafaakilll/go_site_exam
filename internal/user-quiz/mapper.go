package userQuiz

import (
	"time"

	"github.com/mustafaakilll/go-site-exam/db/entity"
)

type BaseRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type BaseResponse struct {
	BaseRequest
	Count int `json:"count"`
}

type CreateUserQuizRequest struct {
	Result   int  `json:"result"`
	IsReview bool `json:"is_review"`

	UserID int `json:"user_id"`
	QuizID int `json:"quiz_id"`
}

type UpdateUserQuizRequest struct {
	ID       int  `json:"id"`
	Result   int  `json:"result"`
	IsReview bool `json:"is_review"`
}

type UserQuizResponseDTO struct {
	Data []UserQuizDTO `json:"data"`
	BaseResponse
}

type UserQuizDTO struct {
	ID           int         `json:"id"`
	Result       int         `json:"result"`
	IsReview     bool        `json:"is_review"`
	StartingTime time.Time   `json:"StartingTime"`
	User         entity.User `json:"user"`
	Quiz         entity.Quiz `json:"quiz"`
}

type Data struct {
	Question QuestionDTO `json:"question"`
}

type QuestionDTO struct {
	ID         int           `json:"id"`
	Text       string        `json:"text"`
	Type       int           `json:"type"`
	Point      int           `json:"point"`
	UserAnswer UserAnswerDTO `json:"user_answer"`
}

type UserAnswerDTO struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	QuizID   int    `json:"quiz_id"`
	AnswerID int    `json:"answer_id"`
	Answer   Answer `json:"answer"`
}

type Question struct {
	ID    int    `json:"ID"`
	Text  string `json:"Text"`
	Type  int    `json:"Type"`
	Point int    `json:"Point"`
}
type Answer struct {
	ID         int      `json:"ID"`
	Text       string   `json:"Text"`
	QuestionID int      `json:"QuestionID"`
	Question   Question `json:"Question"`
}
