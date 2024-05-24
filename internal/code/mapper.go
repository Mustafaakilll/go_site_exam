package code

import "src/github.com/mustafaakilll/go-site-exam/db/entity"

type BaseRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type BaseResponse struct {
	BaseRequest
	Count int `json:"count"`
}

type CreateCodeRequest struct {
	Question string `json:"question"`

	LessonID int `json:"lesson_id"`
}

type UpdateCodeRequest struct {
	ID       int    `json:"id"`
	Question string `json:"question"`
}

type CodeResponseDTO struct {
	Data []CodeDTO `json:"data"`
	BaseResponse
}

type CodeDTO struct {
	ID       int           `json:"id"`
	Question string        `json:"question"`
	Lesson   entity.Lesson `json:"lesson"`
}
