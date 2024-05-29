package codeSubmission

import (
	"github.com/mustafaakilll/go-site-exam/db/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (r *CodeSubmissionRepository) CreateCodeSubmission(codeSubmissionEntity *entity.CodeSubmission) (*entity.CodeSubmission, error) {
	err := r.DB.Create(&codeSubmissionEntity).Error
	return codeSubmissionEntity, err
}

func (r *CodeSubmissionRepository) UpdateCodeSubmission(codeSubmissionEntity entity.CodeSubmission) error {
	err := r.DB.Omit(clause.Associations).Updates(&codeSubmissionEntity).Error
	return err
}

func (r *CodeSubmissionRepository) DeleteCodeSubmission(id int) error {
	err := r.DB.Delete(&entity.CodeSubmission{}, id).Error
	return err
}

func (r *CodeSubmissionRepository) GetCodeSubmissionByID(id int) (*entity.CodeSubmission, error) {
	var codeSubmission entity.CodeSubmission
	err := r.DB.
		Preload("Code").
		Preload("Code.Lesson").
		First(&codeSubmission, id).
		Error
	if err != nil {
		return nil, err
	}
	return &codeSubmission, nil
}

func (r *CodeSubmissionRepository) GetCodeSubmissionsByCodeID(codeID int) ([]entity.CodeSubmission, error) {
	var codeSubmissions []entity.CodeSubmission
	err := r.DB.
		Preload("Code").
		Preload("Code.Lesson").
		Where("code_id = ?", codeID).
		Find(&codeSubmissions).
		Error
	if err != nil {
		return nil, err
	}
	return codeSubmissions, nil
}
