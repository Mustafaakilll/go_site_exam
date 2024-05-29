package question

import (
	"github.com/mustafaakilll/go-site-exam/db/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type QuestionRepository struct {
	DB *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{DB: db}
}

func (r *QuestionRepository) GetQuestions(req *BaseRequest) ([]entity.Question, error) {
	var questions []entity.Question
	query := r.DB
	if req.Limit != 0 {
		query = query.Limit(req.Limit)
	}
	if req.Offset != 0 {
		query = query.Offset(req.Offset)
	}
	err := query.
		Preload("Quiz").
		Find(&questions).
		Error

	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (r *QuestionRepository) CreateQuestion(questionEntity *entity.Question) (*entity.Question, error) {
	err := r.DB.Create(&questionEntity).Error
	return questionEntity, err
}

func (r *QuestionRepository) UpdateQuestion(questionEntity entity.Question) error {
	return r.DB.Omit(clause.Associations).Updates(&questionEntity).Error
}

func (r *QuestionRepository) DeleteQuestion(id int) error {
	return r.DB.Delete(&entity.Question{}, id).Error
}

func (r *QuestionRepository) GetQuestionsByQuizID(req *BaseRequest, quizID int) ([]entity.Question, error) {
	var questions []entity.Question
	query := r.DB
	if req.Limit != 0 {
		query = query.Limit(req.Limit)
	}
	if req.Offset != 0 {
		query = query.Offset(req.Offset)
	}
	err := query.
		Preload("Quiz").
		Where("quiz_id = ?", quizID).
		Find(&questions).
		Error

	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (r *QuestionRepository) GetQuestionByID(id int) (*entity.Question, error) {
	var question entity.Question
	err := r.DB.
		Preload("Quiz").
		First(&question, id).
		Error
	if err != nil {
		return nil, err
	}
	return &question, nil
}

// func (r *QuestionRepository) GetQuestionsWithChoicesByQuizID(quizID int) ([]entity.Question, error) {
// 	var questions []entity.Choice
// 	err := r.DB.
// 		Preload("Question").
// 		Preload("Question.Quiz").
// 		Where("quiz_id = ?", quizID).
// 		Find(&questions).
// 		Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return questions, nil
// }
