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

func (s *UserQuizService) CreateUserQuizzes(userQuizDTO *CreateUserQuizRequest) (*entity.UserQuiz, error) {
	userQuizEntity := new(entity.UserQuiz)
	utils.DTOtoJSON(userQuizDTO, userQuizEntity)

	err := s.repository.CreateUserQuiz(*userQuizEntity)
	if err != nil {
		return nil, err
	}
	return userQuizEntity, nil

}
