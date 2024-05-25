package question

import (
	"github.com/mustafaakilll/go-site-exam/db/entity"

	"gorm.io/gorm"
)

type QuestionRepository struct {
	DB *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{DB: db}
}

func (r *QuestionRepository) GetQuestions(req *BaseRequest) ([]entity.Question, error) {
	var questions []entity.Question
	query := r.DB
	if req.Limit != 0 {
		query = query.Limit(req.Limit)
	}
	if req.Offset != 0 {
		query = query.Offset(req.Offset)
	}
	err := query.
		Preload("Quiz").
		Find(&questions).
		Error

	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (r *QuestionRepository) CreateQuestion(questionEntity entity.Question) error {
	err := r.DB.Create(&questionEntity).Error
	return err
}
