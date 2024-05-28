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
		Preload("Lessons.Teacher").
		Preload("Lessons.Teacher.UserType").
		Find(&users).
		Error
	return users, err
}

func (u *UserRepository) GetUserByID(id int) (entity.User, error) {
	var user entity.User
	err := db.DB.
		Preload("UserType").
		Preload("Lessons").
		Preload("Lessons.Teacher").
		Preload("Lessons.Teacher.UserType").
		First(&user, id).
		Error
	return user, err
}

func (u *UserRepository) CreateUser(user *entity.User) (*entity.User, error) {
	err := db.DB.Create(&user).Error
	return user, err
}

func (u *UserRepository) UpdateUser(user *entity.User) error {
	return db.DB.Updates(&user).Error
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

func (u *UserRepository) SetTeacher(userID, lessonID int) error {
	err := db.DB.
		Model(&entity.User{}).
		Where("id = ?", userID).
		Update("user_type_id", 2).
		Error
	if err != nil {
		return err
	}
	err = db.DB.Model(&entity.Lesson{}).
		Where("id = ?", lessonID).
		Update("teacher_id", userID).
		Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetStudents() ([]entity.User, error) {
	var users []entity.User
	err := db.DB.
		Preload("UserType").
		Preload("Lessons").
		Preload("Lessons.Teacher").
		Preload("Lessons.Teacher.UserType").
		Find(&users, "user_type_id=?", 3).Error
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
		Model(&entity.User{ID: userID}).
		Association("Lessons").
		Append(&entity.Lesson{ID: lessonID})
}

func (u *UserRepository) RemoveLessonFromUser(userID, lessonID int) error {
	return db.DB.
		Model(&entity.User{ID: userID}).
		Association("Lessons").
		Delete(&entity.Lesson{ID: lessonID})
}

func (u *UserRepository) GetStudentsByTeacher(teacherID int) ([]entity.User, error) {
	var users []entity.User
	err := db.DB.
		Preload("UserType").
		Joins("JOIN user_lessons ON user_lessons.user_id = users.id").
		Joins("JOIN lessons ON lessons.id = user_lessons.lesson_id").
		Where("lessons.teacher_id = ?", teacherID).
		Find(&users).
		Error
	return users, err
}

func (u *UserRepository) GetUsersQuizzesByLessonID(lessonID int, req *models.PaginateRequest) ([]entity.User, error) {
	var users []entity.User
	query := db.DB
	if req.Limit != 0 {
		query = query.Limit(req.Limit)
	}
	if req.Offset != 0 {
		query = query.Offset(req.Offset)
	}
	err := query.
		Preload("UserType").
		Joins("JOIN user_lessons ON user_lessons.user_id = users.id").
		Joins("JOIN lessons ON lessons.id = user_lessons.lesson_id").
		Joins("JOIN quizzes ON quizzes.lesson_id = lessons.id").
		Where("users.id not in (select id from user_quizzes where user_quizzes.user_id = users.id)").
		Where("lessons.id = ?", lessonID).
		Where("user_type_id = 3").
		Find(&users).
		Error
	return users, err
}
