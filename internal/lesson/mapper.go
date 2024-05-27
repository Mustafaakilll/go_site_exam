package lesson

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

type UserDTO struct {
	ID        int          `json:"id"`
	FirstName string       `json:"firstname"`
	LastName  string       `json:"lastname"`
	Username  string       `json:"username"`
	Email     string       `json:"email"`
	UserType  *UserTypeDTO `json:"user_types"`
	CreatedAt time.Time    `json:"createdat"`
	UpdatedAt time.Time    `json:"updatedat"`

	Lessons []LessonDTO `json:"lessons"`
}

type UserTypeDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PaginatedUserResponse struct {
	Count int       `json:"count"`
	Data  []UserDTO `json:"data"`
}
