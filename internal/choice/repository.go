package choice

import (
	"src/github.com/mustafaakilll/go-site-exam/db/entity"

	"gorm.io/gorm"
)

type ChoiceRepository struct {
	DB *gorm.DB
}

func NewChoiceRepository(db *gorm.DB) *ChoiceRepository {
	return &ChoiceRepository{DB: db}
}

func (r *ChoiceRepository) GetChoices(req *BaseRequest) ([]entity.Choice, error) {
	var choices []entity.Choice
	query := r.DB
	if req.Limit != 0 {
		query = query.Limit(req.Limit)
	}
	if req.Offset != 0 {
		query = query.Offset(req.Offset)
	}
	err := query.
		Preload("Question").
		Find(&choices).
		Error

	if err != nil {
		return nil, err
	}
	return choices, nil
}

func (r *ChoiceRepository) CreateChoice(choiceEntity entity.Choice) error {
	err := r.DB.Create(&choiceEntity).Error
	return err
}
