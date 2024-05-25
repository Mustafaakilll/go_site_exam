package userAnswer

import (
	"github.com/mustafaakilll/go-site-exam/db/entity"

	"gorm.io/gorm"
)

type UserAnswerRepository struct {
	DB *gorm.DB
}

func NewUserAnswerRepository(db *gorm.DB) *UserAnswerRepository {
	return &UserAnswerRepository{DB: db}
}

func (r *UserAnswerRepository) GetUserAnswers(req *BaseRequest) ([]entity.UserAnswer, error) {
	var userAnswers []entity.UserAnswer
	query := r.DB
	if req.Limit != 0 {
		query = query.Limit(req.Limit)
	}
	if req.Offset != 0 {
		query = query.Offset(req.Offset)
	}
	err := query.Find(&userAnswers).Error

	if err != nil {
		return nil, err
	}
	return userAnswers, nil
}

func (r *UserAnswerRepository) CreateUserAnswer(userAnswerEntity entity.UserAnswer) error {
	err := r.DB.Create(&userAnswerEntity).Error
	return err
}
