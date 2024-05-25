package codeAnswer

import (
	"github.com/mustafaakilll/go-site-exam/db/entity"

	"gorm.io/gorm"
)

type CodeAnswerRepository struct {
	DB *gorm.DB
}

func NewCodeAnswerRepository(db *gorm.DB) *CodeAnswerRepository {
	return &CodeAnswerRepository{DB: db}
}

func (r *CodeAnswerRepository) GetCodeAnswers(req *BaseRequest) ([]entity.CodeAnswer, error) {
	var codeAnswers []entity.CodeAnswer
	query := r.DB
	if req.Limit != 0 {
		query = query.Limit(req.Limit)
	}
	if req.Offset != 0 {
		query = query.Offset(req.Offset)
	}
	err := query.
		Preload("User").
		Preload("Code").
		Find(&codeAnswers).
		Error
	if err != nil {
		return nil, err
	}
	return codeAnswers, nil
}

func (r *CodeAnswerRepository) CreateCodeAnswer(codeAnswerEntity entity.CodeAnswer) error {
	err := r.DB.Create(&codeAnswerEntity).Error
	return err
}
