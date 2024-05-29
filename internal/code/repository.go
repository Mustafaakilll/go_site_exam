package code

import (
	"github.com/mustafaakilll/go-site-exam/db/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (r *CodeRepository) CreateCode(codeEntity *entity.Code) (*entity.Code, error) {
	err := r.DB.Create(&codeEntity).Error
	return codeEntity, err
}

func (r *CodeRepository) UpdateCode(codeEntity entity.Code) error {
	err := r.DB.Omit(clause.Associations).Updates(&codeEntity).Error
	return err
}

func (r *CodeRepository) DeleteCode(id int) error {
	err := r.DB.Delete(&entity.Code{}, id).Error
	return err
}

func (r *CodeRepository) GetCodesByLessonID(req *BaseRequest, lessonID int) ([]entity.Code, error) {
	var codes []entity.Code
	query := r.DB
	if req.Limit != 0 {
		query = query.Limit(req.Limit)
	}
	if req.Offset != 0 {
		query = query.Offset(req.Offset)
	}
	err := query.
		Where("lesson_id = ?", lessonID).
		Find(&codes).
		Error
	if err != nil {
		return nil, err
	}
	return codes, nil
}

func (r *CodeRepository) GetCodeByID(id int) (*entity.Code, error) {
	var code entity.Code
	err := r.DB.Preload("Lesson").First(&code, id).Error
	return &code, err
}

func (r *CodeRepository) GetCodesByTeacherID(req *BaseRequest, teacherID int) ([]entity.Code, error) {
	var codes []entity.Code
	query := r.DB
	if req.Limit != 0 {
		query = query.Limit(req.Limit)
	}
	if req.Offset != 0 {
		query = query.Offset(req.Offset)
	}
	err := query.
		Join("Lesson").
		Where("teacher_id = ?", teacherID).
		Find(&codes).
		Error
	if err != nil {
		return nil, err
	}
	return codes, nil
}
