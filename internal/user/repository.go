package user

import (
	"github.com/mustafaakilll/go-site-exam/db"
	"github.com/mustafaakilll/go-site-exam/db/entity"
	"github.com/mustafaakilll/go-site-exam/pkg/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (u *UserRepository) GetUsers(req *models.PaginateRequest) ([]entity.User, error) {
	var users []entity.User
	query := db.DB
	if req.Limit != 0 {
		query = query.Limit(req.Limit)
	}
	if req.Offset != 0 {
		query = query.Offset(req.Limit)
	}
	err := query.
		Preload("UserType").
		Preload("Lessons").
		Find(&users).
		Error
	return users, err
}

func (u *UserRepository) GetUserByID(id int) (entity.User, error) {
	var user entity.User
	err := db.DB.
		Preload("UserType").
		Preload("Lessons").
		First(&user, id).Error
	return user, err
}

func (u *UserRepository) CreateUser(user *entity.User) (*entity.User, error) {
	err := db.DB.Create(&user).Error
	return user, err
}

func (u *UserRepository) UpdateUser(user *entity.User) error {
	return db.DB.Save(user).Error
}

func (u *UserRepository) DeleteUser(userID int) error {
	return db.DB.Delete(&entity.User{}, userID).Error
}

func (u *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := db.DB.
		Preload("UserType").
		Preload("Lessons").
		First(&user, "email = ?", email).Error
	return &user, err
}

func (u *UserRepository) GetUserByUsername(username string) (entity.User, error) {
	var user entity.User
	err := db.DB.
		Preload("UserType").
		Preload("Lessons").
		First(&user, "username = ?", username).Error
	return user, err
}

func (u *UserRepository) SetTeacher(userID int) error {
	return db.DB.
		Model(&entity.User{}).
		Where("id = ?", userID).
		Update("user_type_id", 2).
		Error
}

func (u *UserRepository) GetStudents() ([]entity.User, error) {
	var users []entity.User
	err := db.DB.
		Preload("UserType").
		Preload("Lessons").
		Find(&users, "user_type_id = ?", 3).Error
	return users, err
}

func (u *UserRepository) GetTeachers() ([]entity.User, error) {
	var users []entity.User
	err := db.DB.
		Preload("UserType").
		Preload("Lessons").
		Find(&users, "user_type_id = ?", 2).Error
	return users, err
}

func (u *UserRepository) AddLessonToUser(userID, lessonID int) error {
	return db.DB.
		Model(&entity.User{}).
		Association("Lessons").
		Append(&entity.Lesson{ID: lessonID}, &entity.User{ID: userID})
}
