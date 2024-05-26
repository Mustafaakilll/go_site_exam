package user

import (
	"errors"

	"github.com/mustafaakilll/go-site-exam/db/entity"
	"github.com/mustafaakilll/go-site-exam/pkg/models"
	"github.com/mustafaakilll/go-site-exam/pkg/utils"
)

type UserService struct {
	repository UserRepository
}

func NewUserService(repository *UserRepository) *UserService {
	return &UserService{repository: *repository}
}

func (u *UserService) GetUsers(req *models.PaginateRequest) (*PaginatedUserResponse, error) {
	users, err := u.repository.GetUsers(req)
	if err != nil {
		return nil, err
	}
	userDTOs := []UserDTO{}
	for i := range users {
		userDTO := new(UserDTO)
		err := utils.JSONtoDTO(users[i], userDTO)

		if err != nil {
			return nil, errors.New("failed to convert user entity to user dto")
		}
		userDTOs = append(userDTOs, *userDTO)

	}
	return &PaginatedUserResponse{
		Count: len(userDTOs),
		Data:  userDTOs,
	}, nil
}

func (u *UserService) GetUserByID(id int) (*UserDTO, error) {
	user, err := u.repository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	userDTO := new(UserDTO)
	utils.JSONtoDTO(user, userDTO)
	return userDTO, nil
}

func (u *UserService) CreateUser(user *CreateUserRequest) (*entity.User, error) {
	userEntity := new(entity.User)
	utils.DTOtoJSON(user, userEntity)

	createdUser, err := u.repository.CreateUser(userEntity)
	if err != nil {
		return nil, err
	}
	return createdUser, nil

}

func (u *UserService) UpdateUser(user *UpdateUserRequest) error {
	userEntity := new(entity.User)
	utils.DTOtoJSON(user, userEntity)

	err := u.repository.UpdateUser(userEntity)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) DeleteUser(userID int) error {
	return u.repository.DeleteUser(userID)
}

func (u *UserService) GetUserByEmail(email string) (*UserDTO, error) {
	user, err := u.repository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	userDTO := new(UserDTO)
	utils.JSONtoDTO(user, userDTO)
	return userDTO, nil
}

func (u *UserService) GetUserByUsername(username string) (*UserDTO, error) {
	user, err := u.repository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	userDTO := new(UserDTO)
	utils.JSONtoDTO(user, userDTO)
	return userDTO, nil
}

func (u *UserService) SetTeacher(userID int) error {
	return u.repository.SetTeacher(userID)
}

func (u *UserService) GetStudents() (*PaginatedUserResponse, error) {
	users, err := u.repository.GetStudents()
	if err != nil {
		return nil, err
	}
	userDTOs := []UserDTO{}
	for i := range users {
		userDTO := new(UserDTO)
		err := utils.JSONtoDTO(users[i], userDTO)
		if err != nil {
			return nil, errors.New("failed to convert user entity to user dto")
		}
		userDTOs = append(userDTOs, *userDTO)
	}
	return &PaginatedUserResponse{
		Count: len(userDTOs),
		Data:  userDTOs,
	}, nil
}

func (u *UserService) AddLessonToUser(userID, lessonID int) error {
	return u.repository.AddLessonToUser(userID, lessonID)
}
