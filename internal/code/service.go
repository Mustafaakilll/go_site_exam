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

func (s *CodeService) CreateCodes(codeDTO *CreateCodeRequest) (*entity.Code, error) {
	codeEntity := new(entity.Code)
	utils.DTOtoJSON(codeDTO, codeEntity)

	err := s.repository.CreateCode(*codeEntity)
	if err != nil {
		return nil, err
	}
	return codeEntity, nil

}
