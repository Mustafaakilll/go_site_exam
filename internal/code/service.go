package code

import (
	"errors"

	"github.com/mustafaakilll/go-site-exam/db/entity"
	"github.com/mustafaakilll/go-site-exam/pkg/utils"
)

type CodeService struct {
	repository CodeRepository
}

func NewCodeService(repository *CodeRepository) *CodeService {
	return &CodeService{repository: *repository}
}

func (s *CodeService) GetCodes(req *BaseRequest) (*CodeResponseDTO, error) {
	codes, err := s.repository.GetCodes(req)
	if err != nil {
		return nil, err
	}
	codeDTOs := []CodeDTO{}
	for i := range codes {
		codeDTO := new(CodeDTO)
		err := utils.JSONtoDTO(codes[i], codeDTO)

		if err != nil {
			return nil, errors.New("failed to convert code entity to code dto")
		}
		codeDTOs = append(codeDTOs, *codeDTO)
	}

	var resultDTO CodeResponseDTO
	resultDTO.Count = len(codeDTOs)
	resultDTO.Data = codeDTOs

	return &resultDTO, nil
}

func (s *CodeService) CreateCode(codeDTO *CreateCodeRequest) (*entity.Code, error) {
	codeEntity := new(entity.Code)
	utils.DTOtoJSON(codeDTO, codeEntity)

	createdCode, err := s.repository.CreateCode(codeEntity)
	if err != nil {
		return nil, err
	}
	return createdCode, nil

}

func (s *CodeService) UpdateCode(codeDTO *UpdateCodeRequest) (*entity.Code, error) {
	codeEntity := new(entity.Code)
	if err := utils.DTOtoJSON(codeDTO, codeEntity); err != nil {
		return nil, err
	}
	err := s.repository.UpdateCode(*codeEntity)
	if err != nil {
		return nil, err
	}
	return codeEntity, nil
}

func (s *CodeService) DeleteCode(id int) error {
	return s.repository.DeleteCode(id)
}

func (s *CodeService) GetCodeByID(id int) (*CodeDTO, error) {
	code, err := s.repository.GetCodeByID(id)
	if err != nil {
		return nil, err
	}
	codeDTO := new(CodeDTO)
	err = utils.JSONtoDTO(code, codeDTO)
	if err != nil {
		return nil, errors.New("failed to convert code entity to code dto")
	}
	return codeDTO, nil
}

func (s *CodeService) GetCodesByLessonID(req *BaseRequest, lessonID int) (*CodeResponseDTO, error) {
	codes, err := s.repository.GetCodesByLessonID(req, lessonID)
	if err != nil {
		return nil, err
	}
	codeDTOs := []CodeDTO{}
	for i := range codes {
		codeDTO := new(CodeDTO)
		err := utils.JSONtoDTO(codes[i], codeDTO)
		if err != nil {
			return nil, errors.New("failed to convert code entity to code dto")
		}
		codeDTOs = append(codeDTOs, *codeDTO)
	}
	var resultDTO CodeResponseDTO
	resultDTO.Count = len(codeDTOs)
	resultDTO.Data = codeDTOs
	return &resultDTO, nil
}

func (s *CodeService) GetCodesByTeacherID(req *BaseRequest, teacherID int) (*CodeResponseDTO, error) {
	codes, err := s.repository.GetCodesByTeacherID(req, teacherID)
	if err != nil {
		return nil, err
	}
	codeDTOs := []CodeDTO{}
	for i := range codes {
		codeDTO := new(CodeDTO)
		err := utils.JSONtoDTO(codes[i], codeDTO)
		if err != nil {
			return nil, errors.New("failed to convert code entity to code dto")
		}
		codeDTOs = append(codeDTOs, *codeDTO)
	}
	var resultDTO CodeResponseDTO
	resultDTO.Count = len(codeDTOs)
	resultDTO.Data = codeDTOs

	return &resultDTO, nil
}
