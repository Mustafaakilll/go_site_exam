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

func (r *LessonRepository) GetLessonByID(id int) (*entity.Lesson, error) {
	var lesson entity.Lesson
	err := r.DB.Preload("Teacher").Preload("Teacher.UserType").First(&lesson, id).Error
	return &lesson, err
}

func (r *LessonRepository) CreateLesson(lessonEntity *entity.Lesson) (*entity.Lesson, error) {
	err := r.DB.Create(&lessonEntity).Error
	return lessonEntity, err
}

func (r *LessonRepository) DeleteLesson(id int) error {
	err := r.DB.Omit(clause.Associations).Delete(&entity.Lesson{}, id).Error
	return err
}

func (r *LessonRepository) UpdateLesson(lessonEntity *entity.Lesson) error {
	err := r.DB.Omit(clause.Associations).Save(&lessonEntity).Error
	return err
}

func (r *LessonRepository) GetLessonById(id int) (*entity.Lesson, error) {
	var lesson entity.Lesson
	err := r.DB.Preload("Teacher").Preload("Teacher.UserType").First(&lesson, id).Error
	return &lesson, err
}

func (r *LessonRepository) GetLessonByTeacher(teacherID int) ([]entity.Lesson, error) {
	var lesson []entity.Lesson
	err := r.DB.
		Preload("Teacher").
		Preload("Teacher.UserType").
		Where("teacher_id = ?", teacherID).
		Find(&lesson).
		Error
	return lesson, err
}

func (r *LessonRepository) GetStudentsByLesson(lessonID int) ([]entity.User, error) {
	var users []entity.User
	err := r.DB.
		Debug().
		Preload("UserType").
		Joins("JOIN user_lessons ON user_lessons.user_id = users.id").
		Where("user_lessons.lesson_id = ?", lessonID).
		Where("user_type_id = 3").
		Find(&users).
		Error
	return users, err
}

func (r *LessonRepository) GetStudentsByNotInLesson(lessonID int) ([]entity.User, error) {
	var users []entity.User
	err := r.DB.
		Preload("UserType").
		Joins("JOIN user_lessons ON user_lessons.user_id <> users.id").
		Where("user_lessons.lesson_id = ?", lessonID).
		Where("user_type_id = 3").
		Find(&users).
		Error
	return users, err
}

func (r *LessonRepository) SetTeacherToLesson(lessonID, userID int) error {
	err := r.DB.
		Model(&entity.Lesson{}).
		Where("id = ?", lessonID).
		Update("teacher_id", userID).
		Error
	return err
}
