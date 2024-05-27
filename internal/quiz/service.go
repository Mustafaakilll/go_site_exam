package quiz

import (
	"errors"

	"github.com/mustafaakilll/go-site-exam/db/entity"
	"github.com/mustafaakilll/go-site-exam/pkg/utils"
)

type QuizService struct {
	repository QuizRepository
}

func NewQuizService(repository *QuizRepository) *QuizService {
	return &QuizService{repository: *repository}
}

func (s *QuizService) GetQuizzes(req *BaseRequest) (*QuizResponseDTO, error) {
	quizzes, err := s.repository.GetQuizzes(req)
	if err != nil {
		return nil, err
	}
	quizDTOs := []QuizDTO{}
	for i := range quizzes {
		quizDTO := new(QuizDTO)
		err := utils.JSONtoDTO(quizzes[i], quizDTO)

		if err != nil {
			return nil, errors.New("failed to convert quiz entity to quiz dto")
		}
		quizDTOs = append(quizDTOs, *quizDTO)
	}

	var resultDTO QuizResponseDTO
	resultDTO.Count = len(quizDTOs)
	resultDTO.Data = quizDTOs

	return &resultDTO, nil
}

func (s *QuizService) CreateQuizzes(quizDTO *CreateQuizRequest) (*entity.Quiz, error) {
	quizEntity := new(entity.Quiz)
	utils.DTOtoJSON(quizDTO, quizEntity)

	err := s.repository.CreateQuiz(*quizEntity)
	if err != nil {
		return nil, err
	}
	return quizEntity, nil

}

func (s *QuizService) UpdateQuiz(quizDTO *UpdateQuizRequest) (*entity.Quiz, error) {
	quizEntity := new(entity.Quiz)
	if err := utils.DTOtoJSON(quizDTO, quizEntity); err != nil {
		return nil, err
	}
	err := s.repository.UpdateQuiz(*quizEntity)
	if err != nil {
		return nil, err
	}
	return quizEntity, nil
}

func (s *QuizService) DeleteQuiz(id int) error {
	return s.repository.DeleteQuiz(id)
}

func (s *QuizService) GetQuizByID(id int) (*QuizResponseDTO, error) {
	quiz, err := s.repository.GetQuizByID(id)
	if err != nil {
		return nil, err
	}
	quizDTO := new(QuizDTO)
	err = utils.JSONtoDTO(quiz, quizDTO)
	if err != nil {
		return nil, errors.New("failed to convert quiz entity to quiz dto")
	}

	var resultDTO QuizResponseDTO
	resultDTO.Count = 1
	resultDTO.Data = []QuizDTO{*quizDTO}

	return &resultDTO, nil
}

func (s *QuizService) GetQuizByTeacher(req *BaseRequest, teacherID int) (*QuizResponseDTO, error) {
	quizzes, err := s.repository.GetQuizByTeacher(req, teacherID)
	if err != nil {
		return nil, err
	}
	quizDTOs := []QuizDTO{}
	for i := range quizzes {
		quizDTO := new(QuizDTO)
		err := utils.JSONtoDTO(quizzes[i], quizDTO)
		if err != nil {
			return nil, errors.New("failed to convert quiz entity to quiz dto")
		}
		quizDTOs = append(quizDTOs, *quizDTO)
	}
	var resultDTO QuizResponseDTO
	resultDTO.Count = len(quizDTOs)
	resultDTO.Data = quizDTOs
	return &resultDTO, nil
}
