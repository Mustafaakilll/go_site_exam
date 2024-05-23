package lesson

import "src/github.com/mustafaakilll/go-site-exam/db/entity"

type BaseRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type BaseResponse struct {
	BaseRequest
	Count int `json:"count"`
}

type CreateLessonRequest struct {
	Name       string `json:"name"`
	Definition string `json:"definition"`
	LessonCode string `json:"lesson_code"`

	TeacherID int `json:"teacher_id"`
}

type UpdateLessonRequest struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Definition string `json:"definition"`
}

type LessonResponseDTO struct {
	Data []LessonDTO `json:"data"`
	BaseResponse
}

type LessonDTO struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	Definition string      `json:"definition"`
	LessonCode string      `json:"lesson_code"`
	Teacher    entity.User `json:"teacher"`
}
