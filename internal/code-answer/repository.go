package codeAnswer

import (
	"github.com/mustafaakilll/go-site-exam/db/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (r *CodeAnswerRepository) CreateCodeAnswer(codeAnswerEntity *entity.CodeAnswer) (*entity.CodeAnswer, error) {
	err := r.DB.Create(&codeAnswerEntity).Error
	return codeAnswerEntity, err
}

func (r *CodeAnswerRepository) UpdateCodeAnswer(codeAnswerEntity entity.CodeAnswer) error {
	err := r.DB.Omit(clause.Associations).Updates(&codeAnswerEntity).Error
	return err
}

func (r *CodeAnswerRepository) DeleteCodeAnswer(id int) error {
	err := r.DB.Delete(&entity.CodeAnswer{}, id).Error
	return err
}
