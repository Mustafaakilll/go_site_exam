package auth

import (
	"github.com/mustafaakilll/go-site-exam/db/entity"
	"github.com/mustafaakilll/go-site-exam/internal/user"
	"github.com/mustafaakilll/go-site-exam/pkg/utils"
)

type AuthService struct {
	repository AuthRepository
	userRepo   user.UserRepository
}

func NewAuthService(repository *AuthRepository, userRepo *user.UserRepository) *AuthService {
	return &AuthService{repository: *repository, userRepo: *userRepo}
}

func (a *AuthService) Login(req *LoginRequest) (*string, error) {
	user, err := a.repository.Login(req)
	if err != nil {
		return nil, err
	}

	if err := utils.CheckPassword(req.Password, user.Password); err != nil {
		return nil, err
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (a *AuthService) RegisterUser(userReq *user.CreateUserRequest) (*string, *entity.User, error) {
	hashedPassword, err := utils.HashPassword(userReq.Password)
	if err != nil {
		return nil, nil, err
	}

	userEntity := new(entity.User)
	if err := utils.DTOtoJSON(userReq, userEntity); err != nil {
		return nil, nil, err
	}
	userEntity.Password = hashedPassword

	createdUser, err := a.userRepo.CreateUser(userEntity)
	if err != nil {
		return nil, nil, err
	}

	token, err := utils.GenerateJWT(createdUser.ID)
	if err != nil {
		return nil, nil, err
	}
	return &token, createdUser, nil
}
