package question

import (
	"errors"

	"github.com/mustafaakilll/go-site-exam/db/entity"
	"github.com/mustafaakilll/go-site-exam/pkg/utils"
)

type QuestionService struct {
	repository QuestionRepository
}

func NewQuestionService(repository *QuestionRepository) *QuestionService {
	return &QuestionService{repository: *repository}
}

func (s *QuestionService) GetQuestions(req *BaseRequest) (*QuestionResponseDTO, error) {
	questions, err := s.repository.GetQuestions(req)
	if err != nil {
		return nil, err
	}
	questionDTOs := []QuestionDTO{}
	for i := range questions {
		questionDTO := new(QuestionDTO)
		err := utils.JSONtoDTO(questions[i], questionDTO)

		if err != nil {
			return nil, errors.New("failed to convert question entity to question dto")
		}
		questionDTOs = append(questionDTOs, *questionDTO)
	}

	var resultDTO QuestionResponseDTO
	resultDTO.Count = len(questionDTOs)
	resultDTO.Data = questionDTOs

	return &resultDTO, nil
}

func (s *QuestionService) CreateQuestion(questionDTO *CreateQuestionRequest) (*entity.Question, error) {
	questionEntity := new(entity.Question)
	utils.DTOtoJSON(questionDTO, questionEntity)

	createdQuestion, err := s.repository.CreateQuestion(questionEntity)
	if err != nil {
		return nil, err
	}
	return createdQuestion, nil

}

func (s *QuestionService) UpdateQuestion(questionDTO *UpdateQuestionRequest) (*entity.Question, error) {
	questionEntity := new(entity.Question)
	if err := utils.DTOtoJSON(questionDTO, questionEntity); err != nil {
		return nil, err
	}
	err := s.repository.UpdateQuestion(*questionEntity)
	if err != nil {
		return nil, err
	}
	return questionEntity, nil
}

func (s *QuestionService) DeleteQuestion(id int) error {
	return s.repository.DeleteQuestion(id)
}

func (s *QuestionService) GetQuestionsByQuizID(req *BaseRequest, quizID int) (*QuestionResponseDTO, error) {
	questions, err := s.repository.GetQuestionsByQuizID(req, quizID)
	if err != nil {
		return nil, err
	}
	questionDTOs := []QuestionDTO{}
	for i := range questions {
		questionDTO := new(QuestionDTO)
		err := utils.JSONtoDTO(questions[i], questionDTO)

		if err != nil {
			return nil, errors.New("failed to convert question entity to question dto")
		}
		questionDTOs = append(questionDTOs, *questionDTO)
	}

	var resultDTO QuestionResponseDTO
	resultDTO.Count = len(questionDTOs)
	resultDTO.Data = questionDTOs

	return &resultDTO, nil
}
