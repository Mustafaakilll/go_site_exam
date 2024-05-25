package codeAnswer

import (
	"errors"

	"github.com/mustafaakilll/go-site-exam/db/entity"
	"github.com/mustafaakilll/go-site-exam/pkg/utils"
)

type CodeAnswerService struct {
	repository CodeAnswerRepository
}

func NewCodeAnswerService(repository *CodeAnswerRepository) *CodeAnswerService {
	return &CodeAnswerService{repository: *repository}
}

func (s *CodeAnswerService) GetCodeAnswers(req *BaseRequest) (*CodeAnswerResponseDTO, error) {
	codeAnswers, err := s.repository.GetCodeAnswers(req)
	if err != nil {
		return nil, err
	}
	codeAnswerDTOs := []CodeAnswerDTO{}
	for i := range codeAnswers {
		codeAnswerDTO := new(CodeAnswerDTO)
		err := utils.JSONtoDTO(codeAnswers[i], codeAnswerDTO)

		if err != nil {
			return nil, errors.New("failed to convert code answer entity to code answer dto")
		}
		codeAnswerDTOs = append(codeAnswerDTOs, *codeAnswerDTO)
	}

	var resultDTO CodeAnswerResponseDTO
	resultDTO.Count = len(codeAnswerDTOs)
	resultDTO.Data = codeAnswerDTOs

	return &resultDTO, nil
}

func (s *CodeAnswerService) CreateCodeAnswers(codeAnswerDTO *CreateCodeAnswerRequest) (*entity.CodeAnswer, error) {
	codeAnswerEntity := new(entity.CodeAnswer)
	utils.DTOtoJSON(codeAnswerDTO, codeAnswerEntity)

	err := s.repository.CreateCodeAnswer(*codeAnswerEntity)
	if err != nil {
		return nil, err
	}
	return codeAnswerEntity, nil
}
