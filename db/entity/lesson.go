package entity

import (
	"time"

	"gorm.io/gorm"
)

type Lesson struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Definition string
	LessonCode string `json:"lesson_code"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	TeacherID int  `json:"teacher_id"`
	Teacher   User `gorm:"foreignKey:TeacherID" json:"teacher"`
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
