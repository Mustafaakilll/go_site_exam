package quiz

import (
	"github.com/mustafaakilll/go-site-exam/db/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type QuizRepository struct {
	DB *gorm.DB
}

func NewQuizRepository(db *gorm.DB) *QuizRepository {
	return &QuizRepository{DB: db}
}

func (r *QuizRepository) GetQuizzes(req *BaseRequest) ([]entity.Quiz, error) {
	var quizzes []entity.Quiz
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
		Preload("Lesson").
		Find(&quizzes).
		Error

	if err != nil {
		return nil, err
	}
	return quizzes, nil
}

func (r *QuizRepository) GetQuizByID(id int) (*entity.Quiz, error) {
	var quiz entity.Quiz
	err := r.DB.
		Preload("Teacher").
		Preload("Teacher.UserType").
		Preload("Lesson").
		First(&quiz, id).
		Error
	if err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (r *QuizRepository) CreateQuiz(quizEntity entity.Quiz) error {
	err := r.DB.Create(&quizEntity).Error
	return err
}

func (r *QuizRepository) UpdateQuiz(quizEntity entity.Quiz) error {
	return r.DB.Omit(clause.Associations).Updates(&quizEntity).Error
}

func (r *QuizRepository) DeleteQuiz(id int) error {
	return r.DB.Omit(clause.Associations).Delete(&entity.Quiz{}, id).Error
}

func (r *QuizRepository) GetQuizByTeacher(req *BaseRequest, teacherID int) ([]entity.Quiz, error) {
	var quizzes []entity.Quiz
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
		Preload("Lesson").
		Where("teacher_id = ?", teacherID).
		Find(&quizzes).
		Error
	return quizzes, err
}

func (r *QuizRepository) GetJoinedUserByQuizID(quizID int) ([]entity.User, error) {
	var users []entity.User
	err := r.DB.
		Preload("UserType").
		Joins("JOIN quiz_user ON users.id = quiz_user.user_id").
		Where("quiz_user.quiz_id = ?", quizID).
		Find(&users).
		Error
	return users, err
}
