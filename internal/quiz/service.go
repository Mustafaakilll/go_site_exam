package quiz

import (
	"errors"
	"src/github.com/mustafaakilll/go-site-exam/db/entity"
	"src/github.com/mustafaakilll/go-site-exam/pkg/utils"
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
