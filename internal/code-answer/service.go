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

	createdCodeAnswer, err := s.repository.CreateCodeAnswer(codeAnswerEntity)
	if err != nil {
		return nil, err
	}
	return createdCodeAnswer, nil
}

func (s *CodeAnswerService) UpdateCodeAnswer(codeAnswerDTO *UpdateCodeAnswerRequest) (*entity.CodeAnswer, error) {
	codeAnswerEntity := new(entity.CodeAnswer)
	if err := utils.DTOtoJSON(codeAnswerDTO, codeAnswerEntity); err != nil {
		return nil, err
	}
	err := s.repository.UpdateCodeAnswer(*codeAnswerEntity)
	if err != nil {
		return nil, err
	}
	return codeAnswerEntity, nil
}

func (s *CodeAnswerService) DeleteCodeAnswer(id int) error {
	return s.repository.DeleteCodeAnswer(id)
}

func (s *CodeAnswerService) GetCodeAnswersByUserID(req *BaseRequest, userID int) (*CodeAnswerResponseDTO, error) {
	codeAnswers, err := s.repository.GetCodeAnswersByUserID(req, userID)
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

func (s *CodeAnswerService) GetCodeAnswerByID(id int) (*CodeAnswerDTO, error) {
	codeAnswer, err := s.repository.GetCodeAnswerByID(id)
	if err != nil {
		return nil, err
	}
	codeAnswerDTO := new(CodeAnswerDTO)
	utils.JSONtoDTO(codeAnswer, codeAnswerDTO)
	return codeAnswerDTO, nil
}
