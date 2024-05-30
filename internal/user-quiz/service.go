package userQuiz

import (
	"errors"

	"github.com/mustafaakilll/go-site-exam/db/entity"
	"github.com/mustafaakilll/go-site-exam/pkg/utils"
)

type UserQuizService struct {
	repository UserQuizRepository
}

func NewUserQuizService(repository *UserQuizRepository) *UserQuizService {
	return &UserQuizService{repository: *repository}
}

func (s *UserQuizService) GetUserQuizzes(req *BaseRequest) (*UserQuizResponseDTO, error) {
	userQuizzes, err := s.repository.GetUserQuizzes(req)
	if err != nil {
		return nil, err
	}
	userQuizDTOs := []UserQuizDTO{}
	for i := range userQuizzes {
		userQuizDTO := new(UserQuizDTO)
		err := utils.JSONtoDTO(userQuizzes[i], userQuizDTO)

		if err != nil {
			return nil, errors.New("failed to convert user quiz entity to user quiz dto")
		}
		userQuizDTOs = append(userQuizDTOs, *userQuizDTO)
	}

	var resultDTO UserQuizResponseDTO
	resultDTO.Count = len(userQuizDTOs)
	resultDTO.Data = userQuizDTOs

	return &resultDTO, nil
}

func (s *UserQuizService) GetUserQuizByID(id int) (*UserQuizDTO, error) {
	userQuiz, err := s.repository.GetUserQuizByID(id)
	if err != nil {
		return nil, err
	}
	userQuizDTO := new(UserQuizDTO)
	err = utils.JSONtoDTO(userQuiz, userQuizDTO)
	if err != nil {
		return nil, errors.New("failed to convert user quiz entity to user quiz dto")
	}
	return userQuizDTO, nil
}

func (s *UserQuizService) CreateUserQuizzes(userQuizDTO *CreateUserQuizRequest) (*entity.UserQuiz, error) {
	userQuizEntity := new(entity.UserQuiz)
	utils.DTOtoJSON(userQuizDTO, userQuizEntity)

	createdUserQuiz, err := s.repository.CreateUserQuiz(userQuizEntity)
	if err != nil {
		return nil, err
	}
	return createdUserQuiz, nil

}

func (s *UserQuizService) UpdateUserQuiz(userQuizDTO *UpdateUserQuizRequest) (*entity.UserQuiz, error) {
	userQuizEntity := new(entity.UserQuiz)
	if err := utils.DTOtoJSON(userQuizDTO, userQuizEntity); err != nil {
		return nil, err
	}
	if err := s.repository.UpdateUserQuiz(*userQuizEntity); err != nil {
		return nil, err
	}
	return userQuizEntity, nil
}

func (s *UserQuizService) DeleteUserQuiz(id int) error {
	return s.repository.DeleteUserQuiz(id)
}

func (s *UserQuizService) GetUsersQuizzesByUserID(request *BaseRequest, userID int) (*UserQuizResponseDTO, error) {
	userQuizzes, err := s.repository.GetUserQuizByUserID(userID)
	if err != nil {
		return nil, err
	}
	userQuizDTOs := []UserQuizDTO{}
	for i := range userQuizzes {
		userQuizDTO := new(UserQuizDTO)
		err := utils.JSONtoDTO(userQuizzes[i], userQuizDTO)
		if err != nil {
			return nil, errors.New("failed to convert user quiz entity to user quiz dto")
		}
		userQuizDTOs = append(userQuizDTOs, *userQuizDTO)
	}
	var resultDTO UserQuizResponseDTO
	resultDTO.Count = len(userQuizDTOs)
	resultDTO.Data = userQuizDTOs
	return &resultDTO, nil
}

func (s *UserQuizService) GetUsersQuizzesByLessonID(request *BaseRequest, lessonID int) (*UserQuizResponseDTO, error) {
	userQuizzes, err := s.repository.GetUsersQuizzesByLessonID(lessonID, request)
	if err != nil {
		return nil, err
	}
	userQuizDTOs := []UserQuizDTO{}
	for i := range userQuizzes {
		userQuizDTO := new(UserQuizDTO)
		err := utils.JSONtoDTO(userQuizzes[i], userQuizDTO)
		if err != nil {
			return nil, errors.New("failed to convert user quiz entity to user quiz dto")
		}
		userQuizDTOs = append(userQuizDTOs, *userQuizDTO)
	}
	var resultDTO UserQuizResponseDTO
	resultDTO.Count = len(userQuizDTOs)
	resultDTO.Data = userQuizDTOs
	return &resultDTO, nil
}

func (s *UserQuizService) GetUserQuizWithAnswersByUserAndQuizID(userID, quizID int) ([]Data, error) {
	quizzes, err := s.repository.GetUserQuizByUserID(userID)
	if err != nil {
		return nil, err
	}
	data := []Data{}
	for _, quiz := range quizzes {
		if quiz.QuizID == quizID {
			userQuizDTO := new(UserQuizDTO)
			err := utils.JSONtoDTO(quiz, userQuizDTO)
			if err != nil {
				return nil, errors.New("failed to convert user quiz entity to user quiz dto")
			}
			questions, _ := s.repository.GetQuestionsByQuizID(quizID)
			for _, question := range questions {
				userAnswers, _ := s.repository.GetUsersAnswersByQuestionID(question.ID)
				if len(userAnswers) > 0 {
					for _, userAnswer := range userAnswers {
						if userAnswer.Answer.QuestionID != question.ID {
							continue
						}
						answer, _ := s.repository.GetAnswerByAnswerID(userAnswer.AnswerID)

						answerDTO := &Answer{
							ID:         answer.ID,
							Text:       answer.Text,
							QuestionID: answer.QuestionID,
						}
						userAnswerDTO := &UserAnswerDTO{
							ID:       userAnswer.ID,
							UserID:   userAnswer.UserID,
							QuizID:   userAnswer.QuizID,
							AnswerID: userAnswer.AnswerID,
							Answer:   *answerDTO,
						}
						questionDTO := &QuestionDTO{
							ID:         question.ID,
							Text:       question.Text,
							Type:       question.Type,
							Point:      question.Point,
							UserAnswer: *userAnswerDTO,
						}
						data = append(data, Data{Question: *questionDTO})
						answerDTO = nil
						userAnswerDTO = nil
						questionDTO = nil
					}
				}
			}
		}
	}

	return data, nil
}

func (s *UserQuizService) GetUserQuizByUserAndQuizID(userID, quizID int) (*UserQuizDTO, error) {
	userQuiz, err := s.repository.GetUserQuizByUserAndQuizID(userID, quizID)
	if err != nil {
		return nil, err
	}
	userQuizDTO := new(UserQuizDTO)
	err = utils.JSONtoDTO(userQuiz, userQuizDTO)
	if err != nil {
		return nil, errors.New("failed to convert user quiz entity to user quiz dto")
	}
	return userQuizDTO, nil
}
