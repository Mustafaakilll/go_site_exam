package lesson

import (
	"errors"

	"github.com/mustafaakilll/go-site-exam/db/entity"
	"github.com/mustafaakilll/go-site-exam/pkg/utils"
)

type LessonService struct {
	repository LessonRepository
}

func NewLessonService(repository *LessonRepository) *LessonService {
	return &LessonService{repository: *repository}
}

func (s *LessonService) GetLessons(req *BaseRequest) (*LessonResponseDTO, error) {
	lessons, err := s.repository.GetLessons(req)
	if err != nil {
		return nil, err
	}
	lessonDTOs := []LessonDTO{}
	for i := range lessons {
		lessonDTO := new(LessonDTO)
		err := utils.JSONtoDTO(lessons[i], lessonDTO)

		if err != nil {
			return nil, errors.New("failed to convert lesson entity to lesson dto")
		}
		lessonDTOs = append(lessonDTOs, *lessonDTO)
	}

	var resultDTO LessonResponseDTO
	resultDTO.Count = len(lessonDTOs)
	resultDTO.Data = lessonDTOs

	return &resultDTO, nil
}

func (s *LessonService) CreateLessons(lessonDTO *CreateLessonRequest) (*entity.Lesson, error) {
	lessonEntity := new(entity.Lesson)
	utils.DTOtoJSON(lessonDTO, lessonEntity)

	err := s.repository.CreateLesson(*lessonEntity)
	if err != nil {
		return nil, err
	}
	return lessonEntity, nil

}
