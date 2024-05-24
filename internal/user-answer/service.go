package user_answer

import (
	"errors"
	"src/github.com/mustafaakilll/go-site-exam/db/entity"
	"src/github.com/mustafaakilll/go-site-exam/pkg/utils"
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

func (s *UserAnswerService) CreateUserAnswers(userAnswerDTO *CreateUserAnswerRequest) (*entity.UserAnswer, error) {
	userAnswerEntity := new(entity.UserAnswer)
	utils.DTOtoJSON(userAnswerDTO, userAnswerEntity)

	err := s.repository.CreateUserAnswer(*userAnswerEntity)
	if err != nil {
		return nil, err
	}
	return userAnswerEntity, nil

}
