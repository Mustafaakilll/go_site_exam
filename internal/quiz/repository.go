package quiz

import (
	"src/github.com/mustafaakilll/go-site-exam/db/entity"

	"gorm.io/gorm"
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

func (r *QuizRepository) CreateQuiz(quizEntity entity.Quiz) error {
	err := r.DB.Create(&quizEntity).Error
	return err
}
