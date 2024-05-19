package entity

import (
	"time"

	"gorm.io/gorm"
)

type Lesson struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Definition string
	LessonCode string
	CreatedAt  time.Time
	UpdatedAt  time.Time

	TeacherID int
	Teacher   User `gorm:"foreignKey:TeacherID"`
}

func (u *Lesson) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}

func (u *Lesson) BeforeSave(tx *gorm.DB) (err error) {
	if u.ID == 0 {
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()
	return
}
