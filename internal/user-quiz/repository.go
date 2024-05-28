package userQuiz

import (
	"github.com/mustafaakilll/go-site-exam/db/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (r *UserQuizRepository) CreateUserQuiz(userQuizEntity *entity.UserQuiz) (*entity.UserQuiz, error) {
	err := r.DB.Create(&userQuizEntity).Error
	return userQuizEntity, err
}

func (r *UserQuizRepository) UpdateUserQuiz(userQuizEntity entity.UserQuiz) error {
	err := r.DB.Omit(clause.Associations).Updates(&userQuizEntity).Error
	return err
}

func (r *UserQuizRepository) DeleteUserQuiz(id int) error {
	err := r.DB.Omit(clause.Associations).Delete(&entity.UserQuiz{}, id).Error
	return err
}

func (r *UserQuizRepository) GetUserQuizByID(id uint) (entity.UserQuiz, error) {
	var userQuiz entity.UserQuiz
	err := r.DB.
		Preload("User").
		Preload("Quiz").
		Where("id = ?", id).
		First(&userQuiz).
		Error
	return userQuiz, err
}

func (r *UserQuizRepository) GetUserQuizByUserIDAndQuizID(userID, quizID uint) (entity.UserQuiz, error) {
	var userQuiz entity.UserQuiz
	err := r.DB.
		Preload("User").
		Preload("Quiz").
		Where("user_id = ? AND quiz_id = ?", userID, quizID).
		First(&userQuiz).
		Error
	return userQuiz, err
}

func (r *UserQuizRepository) GetUserQuizByUserID(userID uint) ([]entity.UserQuiz, error) {
	var userQuizzes []entity.UserQuiz
	err := r.DB.
		Preload("User").
		Preload("Quiz").
		Where("user_id = ?", userID).
		Find(&userQuizzes).
		Error
	return userQuizzes, err
}

func (r *UserQuizRepository) GetUserQuizByQuizID(quizID uint) ([]entity.UserQuiz, error) {
	var userQuizzes []entity.UserQuiz
	err := r.DB.
		Preload("User").
		Preload("Quiz").
		Where("quiz_id = ?", quizID).
		Find(&userQuizzes).
		Error
	return userQuizzes, err
}

func (r *UserQuizRepository) GetUsersQuizzesByUserID(userID int, req *BaseRequest) ([]entity.UserQuiz, error) {
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
		Where("user_id = ?", userID).
		Find(&userQuizzes).
		Error
	return userQuizzes, err
}
