package codeSubmission

import (
	"github.com/mustafaakilll/go-site-exam/db/entity"

	"gorm.io/gorm"
)

type CodeSubmissionRepository struct {
	DB *gorm.DB
}

func NewCodeSubmissionRepository(db *gorm.DB) *CodeSubmissionRepository {
	return &CodeSubmissionRepository{DB: db}
}

func (r *CodeSubmissionRepository) GetCodeSubmissions(req *BaseRequest) ([]entity.CodeSubmission, error) {
	var codeSubmissions []entity.CodeSubmission
	query := r.DB
	if req.Limit != 0 {
		query = query.Limit(req.Limit)
	}
	if req.Offset != 0 {
		query = query.Offset(req.Offset)
	}
	err := query.
		Preload("Code").
		Find(&codeSubmissions).
		Error
	if err != nil {
		return nil, err
	}
	return codeSubmissions, nil
}

func (r *CodeSubmissionRepository) CreateCodeSubmission(codeSubmissionEntity entity.CodeSubmission) error {
	err := r.DB.Create(&codeSubmissionEntity).Error
	return err
}
