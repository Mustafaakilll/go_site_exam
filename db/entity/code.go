package entity

import (
	"time"

	"gorm.io/gorm"
)

type Code struct {
	ID        int `gorm:"primaryKey"`
	Question  string
	CreatedAt time.Time
	UpdatedAt time.Time

	LessonID int    `json:"lesson_id"`
	Lesson   Lesson `gorm:"foreignKey:LessonID" json:"lesson"`
}

func (u *Code) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}

func (u *Code) BeforeSave(tx *gorm.DB) (err error) {
	if u.ID == 0 {
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()
	return
}
