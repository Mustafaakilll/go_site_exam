package answer

import (
	"errors"

	"github.com/mustafaakilll/go-site-exam/db/entity"
	"github.com/mustafaakilll/go-site-exam/pkg/utils"
)

type AnswerService struct {
	repository AnswerRepository
}

func NewAnswerService(repository *AnswerRepository) *AnswerService {
	return &AnswerService{repository: *repository}
}

func (s *AnswerService) GetAnswers(req *BaseRequest) (*AnswerResponseDTO, error) {
	answers, err := s.repository.GetAnswers(req)
	if err != nil {
		return nil, err
	}
	answerDTOs := []AnswerDTO{}
	for i := range answers {
		answerDTO := new(AnswerDTO)
		err := utils.JSONtoDTO(answers[i], answerDTO)

		if err != nil {
			return nil, errors.New("failed to convert answer entity to answer dto")
		}
		answerDTOs = append(answerDTOs, *answerDTO)
	}

	var resultDTO AnswerResponseDTO
	resultDTO.Count = len(answerDTOs)
	resultDTO.Data = answerDTOs

	return &resultDTO, nil
}

func (s *AnswerService) CreateAnswers(answerDTO *CreateAnswerRequest) (*entity.Answer, error) {
	answerEntity := new(entity.Answer)
	utils.DTOtoJSON(answerDTO, answerEntity)

	err := s.repository.CreateAnswer(*answerEntity)
	if err != nil {
		return nil, err
	}
	return answerEntity, nil
}

func (s *AnswerService) UpdateAnswers(answerDTO *UpdateAnswerRequest) (*entity.Answer, error) {
	answerEntity := new(entity.Answer)
	if err := utils.DTOtoJSON(answerDTO, answerEntity); err != nil {
		return nil, err
	}

	err := s.repository.UpdateAnswer(*answerEntity)
	if err != nil {
		return nil, err
	}
	return answerEntity, nil
}

func (s *AnswerService) DeleteAnswer(id int) error {
	return s.repository.DeleteAnswer(id)
}
