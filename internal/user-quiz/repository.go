package userQuiz

import (
	"time"

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
	userQuizEntity.StartingTime = time.Now()
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

func (r *UserQuizRepository) GetUserQuizByID(id int) (entity.UserQuiz, error) {
	var userQuiz entity.UserQuiz
	err := r.DB.
		Preload("User").
		Preload("Quiz").
		Where("id = ?", id).
		First(&userQuiz).
		Error
	return userQuiz, err
}

// func (r *UserQuizRepository) GetUserQuizByUserIDAndQuizID(userID, quizID uint) (entity.UserQuiz, error) {
// 	var userQuiz entity.UserQuiz
// 	err := r.DB.
// 		Preload("User").
// 		Preload("Quiz").
// 		Where("user_id = ? AND quiz_id = ?", userID, quizID).
// 		First(&userQuiz).
// 		Error
// 	return userQuiz, err
// }

func (r *UserQuizRepository) GetUserQuizByUserID(userID int) ([]entity.UserQuiz, error) {
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

func (r *UserQuizRepository) GetUsersQuizzesByLessonID(lessonID int, req *BaseRequest) ([]entity.UserQuiz, error) {
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
		Preload("Quiz.Lesson").
		Where("lesson_id = ?", lessonID).
		Find(&userQuizzes).
		Error
	return userQuizzes, err
}

// func (r *UserQuizRepository) GetUserQuizWithAnswersByUserAndQuizID(userID, quizID int) (entity.UserQuiz, error) {
// 	var userQuiz entity.UserQuiz
// 	// r.DB.Exec("SELECT * FROM user_quizzes WHERE user_id = ? AND quiz_id = ?", userID, quizID)
// 	// select * from user_answers join user_quizzes on user_quizzes.quiz_id = user_answers.quiz_id join questions on questions.quiz_id = user_answers.quiz_id;
// 	err := r.DB.
// 		Joins("JOIN user_answers ON user_quizzes.quiz_id = user_answers.quiz_id").
// 		Joins("JOIN questions ON questions.quiz_id = user_answers.quiz_id").
// 		Where("user_quizzes.quiz_id = ? and user_quizzes.user_id = ?", quizID, userID).
// 		Find(&userQuiz).Error
//
// 	fmt.Printf("%+v", userQuiz)
//
// 	// err := r.DB.
// 	// 	Preload("User").
// 	// 	Preload("User.UserType").
// 	// 	Preload("Quiz").
// 	// 	Preload("Quiz.Lesson").
// 	// 	Preload("Quiz.Teacher").
// 	// 	Joins("LEFT JOIN user_answers ON user_quizzes.id = user_answers.user_id").
// 	// 	Joins("LEFT JOIN questions ON questions.quiz_id = user_quizzes.quiz_id").
// 	// 	Where("user_quizzes.quiz_id = ? and user_quizzes.user_id = ?", quizID, userID).
// 	// 	Find(&userQuiz).Error
// 	return userQuiz, err
// }

// func (r *UserQuizRepository) GetUsersQuizzesByUserID(userID, quizID int) (entity.UserQuiz, error) {
// 	var userQuiz entity.UserQuiz
// 	err := r.DB.
// 		Preload("User").
// 		Preload("Quiz").
// 		Preload("Quiz.Lesson").
// 		Preload("Quiz.Teacher").
// 		Preload("UserAnswers").
// 		Preload("UserAnswers.Question").
// 		Preload("UserAnswers.Question.Answers").
// 		Where("quiz_id = ? AND user_id = ?", quizID, userID).
// 		First(&userQuiz).
// 		Error
// 	return userQuiz, err
// }

func (r *UserQuizRepository) GetQuestionsByQuizID(quizID int) ([]entity.Question, error) {
	var questions []entity.Question
	err := r.DB.
		Preload("Quiz").
		Where("quiz_id = ?", quizID).
		Find(&questions).
		Error
	return questions, err
}

func (r *UserQuizRepository) GetUsersAnswersByQuestionID(questionID int) ([]entity.UserAnswer, error) {
	var userAnswers []entity.UserAnswer
	err := r.DB.
		Preload("Answer", "question_id=?", questionID).
		Preload("Answer.Question").
		Find(&userAnswers).
		Error
	return userAnswers, err
}

func (r *UserQuizRepository) GetAnswerByAnswerID(answerID int) (entity.Answer, error) {
	var answer entity.Answer
	err := r.DB.
		Preload("Question").
		Where("id = ?", answerID).
		First(&answer).
		Error
	return answer, err
}

func (r *UserQuizRepository) GetUserQuizByUserAndQuizID(userID, quizID int) (entity.UserQuiz, error) {
	var answer entity.UserQuiz
	err := r.DB.
		Where("user_id = ? and quiz_id=?", userID, quizID).
		First(&answer).
		Error
	return answer, err
}
