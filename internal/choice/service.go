package choice

import (
	"errors"
	"src/github.com/mustafaakilll/go-site-exam/db/entity"
	"src/github.com/mustafaakilll/go-site-exam/pkg/utils"
)

type ChoiceService struct {
	repository ChoiceRepository
}

func NewChoiceService(repository *ChoiceRepository) *ChoiceService {
	return &ChoiceService{repository: *repository}
}

func (s *ChoiceService) GetChoices(req *BaseRequest) (*ChoiceResponseDTO, error) {
	choices, err := s.repository.GetChoices(req)
	if err != nil {
		return nil, err
	}
	choiceDTOs := []ChoiceDTO{}
	for i := range choices {
		choiceDTO := new(ChoiceDTO)
		err := utils.JSONtoDTO(choices[i], choiceDTO)

		if err != nil {
			return nil, errors.New("failed to convert choice entity to choice dto")
		}
		choiceDTOs = append(choiceDTOs, *choiceDTO)
	}

	var resultDTO ChoiceResponseDTO
	resultDTO.Count = len(choiceDTOs)
	resultDTO.Data = choiceDTOs

	return &resultDTO, nil
}

func (s *ChoiceService) CreateChoices(choiceDTO *CreateChoiceRequest) (*entity.Choice, error) {
	choiceEntity := new(entity.Choice)
	utils.DTOtoJSON(choiceDTO, choiceEntity)

	err := s.repository.CreateChoice(*choiceEntity)
	if err != nil {
		return nil, err
	}
	return choiceEntity, nil

}
