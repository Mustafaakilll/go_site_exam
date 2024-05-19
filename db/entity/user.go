package entity

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        int `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time

	UserTypeID int      // 0: admin, 1: teacher, 2: student
	UserType   UserType `json:"user_types" gorm:"foreignKey:UserTypeID;"`
	Lessons    []Lesson `gorm:"many2many:user_lessons;"`
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.ID == 0 {
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()
	u.UserTypeID = 3
	u.Username = u.FirstName + u.LastName + strings.Split(uuid.New().String(), "-")[0]
	return
}
