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

func (s *LessonService) GetLessonByID(id int) (*LessonDTO, error) {
	lesson, err := s.repository.GetLessonByID(id)
	if err != nil {
		return nil, err
	}
	lessonDTO := new(LessonDTO)
	err = utils.JSONtoDTO(lesson, lessonDTO)
	if err != nil {
		return nil, errors.New("failed to convert lesson entity to lesson dto")
	}
	return lessonDTO, nil
}

func (s *LessonService) CreateLessons(lessonDTO *CreateLessonRequest) (*entity.Lesson, error) {
	lessonEntity := new(entity.Lesson)
	utils.DTOtoJSON(lessonDTO, lessonEntity)

	createdLesson, err := s.repository.CreateLesson(lessonEntity)
	if err != nil {
		return nil, err
	}
	return createdLesson, nil
}

func (s *LessonService) DeleteLesson(id int) error {
	err := s.repository.DeleteLesson(id)
	return err
}

func (s *LessonService) UpdateLesson(lessonDTO *UpdateLessonRequest) (*entity.Lesson, error) {
	lessonEntity := new(entity.Lesson)
	if err := utils.DTOtoJSON(lessonDTO, lessonEntity); err != nil {
		return nil, err
	}

	err := s.repository.UpdateLesson(lessonEntity)
	if err != nil {
		return nil, err
	}
	return lessonEntity, nil
}

func (s *LessonService) GetLessonByTeacher(teacherID int) (*LessonResponseDTO, error) {
	lessons, err := s.repository.GetLessonByTeacher(teacherID)
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

func (s *LessonService) GetStudentsByLesson(lessonID int) (*PaginatedUserResponse, error) {
	users, err := s.repository.GetStudentsByLesson(lessonID)

	if err != nil {
		return nil, err
	}
	userDTOs := []UserDTO{}
	for i := range users {
		userDTO := new(UserDTO)
		err := utils.JSONtoDTO(users[i], userDTO)

		if err != nil {
			return nil, errors.New("failed to convert lesson entity to lesson dto")
		}
		userDTOs = append(userDTOs, *userDTO)
	}

	var resultDTO PaginatedUserResponse
	resultDTO.Count = len(userDTOs)
	resultDTO.Data = userDTOs

	return &resultDTO, nil
}

func (s *LessonService) GetStudentsByNotInLesson(lessonID int) (*PaginatedUserResponse, error) {
	users, err := s.repository.GetStudentsByNotInLesson(lessonID)

	if err != nil {
		return nil, err
	}
	userDTOs := []UserDTO{}
	for i := range users {
		userDTO := new(UserDTO)
		err := utils.JSONtoDTO(users[i], userDTO)

		if err != nil {
			return nil, errors.New("failed to convert lesson entity to lesson dto")
		}
		userDTOs = append(userDTOs, *userDTO)
	}

	var resultDTO PaginatedUserResponse
	resultDTO.Count = len(userDTOs)
	resultDTO.Data = userDTOs

	return &resultDTO, nil
}

func (s *LessonService) SetTeacherToLesson(lessonID, userID int) error {
	err := s.repository.SetTeacherToLesson(lessonID, userID)
	if err != nil {
		return err
	}
	return nil
}
