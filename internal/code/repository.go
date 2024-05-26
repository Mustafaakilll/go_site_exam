package code

import (
	"github.com/mustafaakilll/go-site-exam/db/entity"

	"gorm.io/gorm"
)

type CodeRepository struct {
	DB *gorm.DB
}

func NewCodeRepository(db *gorm.DB) *CodeRepository {
	return &CodeRepository{DB: db}
}

func (r *CodeRepository) GetCodes(req *BaseRequest) ([]entity.Code, error) {
	var codes []entity.Code
	query := r.DB
	if req.Limit != 0 {
		query = query.Limit(req.Limit)
	}
	if req.Offset != 0 {
		query = query.Offset(req.Offset)
	}
	err := query.
		Preload("Lesson").
		Find(&codes).
		Error

	if err != nil {
		return nil, err
	}
	return codes, nil
}

func (r *CodeRepository) CreateCode(codeEntity entity.Code) error {
	err := r.DB.Create(&codeEntity).Error
	return err
}

func (r *CodeRepository) UpdateCode(codeEntity entity.Code) error {
	err := r.DB.Save(&codeEntity).Error
	return err
}

func (r *CodeRepository) DeleteCode(id int) error {
	err := r.DB.Delete(&entity.Code{}, id).Error
	return err
}
