package lesson

import (
	"github.com/mustafaakilll/go-site-exam/db/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (r *LessonRepository) DeleteLesson(id int) error {
	err := r.DB.Omit(clause.Associations).Delete(&entity.Lesson{}, id).Error
	return err
}

func (r *LessonRepository) UpdateLesson(lessonEntity entity.Lesson) error {
	err := r.DB.Omit(clause.Associations).Save(&lessonEntity).Error
	return err
}
