package auth

import (
	"github.com/mustafaakilll/go-site-exam/db/entity"
	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (a *AuthRepository) Login(req *LoginRequest) (*entity.User, error) {
	var user entity.User
	err := a.DB.
		Model(&entity.User{}).
		Where("username = ?", req.Username).
		First(&user).Error

	return &user, err
}
