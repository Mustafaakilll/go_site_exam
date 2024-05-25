package userQuiz

import (
	"github.com/mustafaakilll/go-site-exam/db/entity"

	"gorm.io/gorm"
)

type UserQuizRepository struct {
	DB *gorm.DB
}

func NewUserQuizRepository(db *gorm.DB) *UserQuizRepository {
	return &UserQuizRepository{DB: db}
}

func (r *UserQuizRepository) GetUserQuizzes(req *BaseRequest) ([]entity.UserQuiz, error) {
	var userQuizzes []entity.UserQuiz
	query := r.DB
	if req.Limit != 0 {
		query = query.Limit(req.Limit)
	}
	if req.Offset != 0 {
		query = query.Offset(req.Offset)
	}
	err := query.
		Preload("User").
		Preload("Quiz").
		Find(&userQuizzes).
		Error

	if err != nil {
		return nil, err
	}
	return userQuizzes, nil
}

func (r *UserQuizRepository) CreateUserQuiz(userQuizEntity entity.UserQuiz) error {
	err := r.DB.Create(&userQuizEntity).Error
	return err
}
