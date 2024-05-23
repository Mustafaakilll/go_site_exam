package lesson

import (
	"src/github.com/mustafaakilll/go-site-exam/db/entity"

	"gorm.io/gorm"
)

type LessonRepository struct {
	DB *gorm.DB
}

func NewLessonRepository(db *gorm.DB) *LessonRepository {
	return &LessonRepository{DB: db}
}

func (r *LessonRepository) GetLessons(req *BaseRequest) ([]entity.Lesson, error) {
	var lessons []entity.Lesson
	query := r.DB
	if req.Limit != 0 {
		query = query.Limit(req.Limit)
	}
	if req.Offset != 0 {
		query = query.Offset(req.Offset)
	}
	err := query.
		Preload("Teacher").
		Preload("Teacher.UserType").
		Find(&lessons).
		Error

	if err != nil {
		return nil, err
	}
	return lessons, nil
}

func (r *LessonRepository) CreateLesson(lessonEntity entity.Lesson) error {
	err := r.DB.Create(&lessonEntity).Error
	return err
}
