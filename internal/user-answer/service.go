package userAnswer

import (
	"errors"

	"github.com/mustafaakilll/go-site-exam/db/entity"
	"github.com/mustafaakilll/go-site-exam/pkg/utils"
)

type UserAnswerService struct {
	repository UserAnswerRepository
}

func NewUserAnswerService(repository *UserAnswerRepository) *UserAnswerService {
	return &UserAnswerService{repository: *repository}
}

func (s *UserAnswerService) GetUserAnswers(req *BaseRequest) (*UserAnswerResponseDTO, error) {
	userAnswers, err := s.repository.GetUserAnswers(req)
	if err != nil {
		return nil, err
	}
	userAnswerDTOs := []UserAnswerDTO{}
	for i := range userAnswers {
		userAnswerDTO := new(UserAnswerDTO)
		err := utils.JSONtoDTO(userAnswers[i], userAnswerDTO)

		if err != nil {
			return nil, errors.New("failed to convert user answer entity to user answer dto")
		}
		userAnswerDTOs = append(userAnswerDTOs, *userAnswerDTO)
	}

	var resultDTO UserAnswerResponseDTO
	resultDTO.Count = len(userAnswerDTOs)
	resultDTO.Data = userAnswerDTOs

	return &resultDTO, nil
}

func (s *UserAnswerService) CreateUserAnswer(userAnswerDTO *CreateUserAnswerRequest) (*entity.UserAnswer, error) {
	userAnswerEntity := new(entity.UserAnswer)
	utils.DTOtoJSON(userAnswerDTO, userAnswerEntity)

	createdUserAnswer, err := s.repository.CreateUserAnswer(userAnswerEntity)
	if err != nil {
		return nil, err
	}
	return createdUserAnswer, nil

}

func (s *UserAnswerService) UpdateUserAnswer(userAnswerDTO *UpdateUserAnswerRequest) (*entity.UserAnswer, error) {
	userAnswerEntity := new(entity.UserAnswer)
	if err := utils.DTOtoJSON(userAnswerDTO, userAnswerEntity); err != nil {
		return nil, err
	}
	if err := s.repository.UpdateUserAnswer(*userAnswerEntity); err != nil {
		return nil, err
	}
	return userAnswerEntity, nil
}

func (s *UserAnswerService) DeleteUserAnswer(id int) error {
	return s.repository.DeleteUserAnswer(id)
}

func (s *UserAnswerService) GetUserAnswerByQuestionID(id int) (*UserAnswerDTO, error) {
	userAnswer, err := s.repository.GetUserAnswerByQuestionID(id)
	if err != nil {
		return nil, err
	}
	userAnswerDTO := new(UserAnswerDTO)
	err = utils.JSONtoDTO(userAnswer, userAnswerDTO)
	if err != nil {
		return nil, errors.New("failed to convert user answer entity to user answer dto")
	}
	return userAnswerDTO, nil
}
