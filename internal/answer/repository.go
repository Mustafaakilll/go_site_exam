package answer

import (
	"github.com/mustafaakilll/go-site-exam/db/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AnswerRepository struct {
	DB *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) *AnswerRepository {
	return &AnswerRepository{DB: db}
}

func (r *AnswerRepository) GetAnswers(req *BaseRequest) ([]entity.Answer, error) {
	var answers []entity.Answer
	query := r.DB
	if req.Limit != 0 {
		query = query.Limit(req.Limit)
	}
	if req.Offset != 0 {
		query = query.Offset(req.Offset)
	}
	err := query.
		Preload("Choice").
		Preload("Question").
		Preload("User").
		Find(&answers).
		Error

	if err != nil {
		return nil, err
	}
	return answers, nil
}

func (r *AnswerRepository) CreateAnswer(answerEntity *entity.Answer) (*entity.Answer, error) {
	err := r.DB.Create(&answerEntity).Error
	return answerEntity, err
}

func (r *AnswerRepository) UpdateAnswer(answerEntity entity.Answer) error {
	err := r.DB.Omit(clause.Associations).Save(&answerEntity).Error
	return err
}

func (r *AnswerRepository) DeleteAnswer(id int) error {
	err := r.DB.Delete(&entity.Answer{}, id).Error
	return err
}
