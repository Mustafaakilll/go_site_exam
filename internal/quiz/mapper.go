package quiz

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

type CreateQuizRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Duration    int       `json:"duration"`
	StartTime   time.Time `json:"start_time"`

	LessonID int `json:"lesson_id"`
}

type UpdateQuizRequest struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Duration    int       `json:"duration"`
	StartTime   time.Time `json:"start_time"`
}

type QuizResponseDTO struct {
	Data []QuizDTO `json:"data"`
	BaseResponse
}

type QuizDTO struct {
	ID          int           `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Duration    int           `json:"duration"`
	StartTime   time.Time     `json:"start_time"`
	EndTime     time.Time     `json:"end_time"`
	Teacher     entity.User   `json:"teacher"`
	Lesson      entity.Lesson `json:"lesson"`
}
