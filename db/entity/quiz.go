package entity

import (
	"time"

	"gorm.io/gorm"
)

type Quiz struct {
	ID          int `gorm:"primaryKey"`
	Title       string
	Description string
	Duration    int
	StartTime   time.Time
	EndTime     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time

	TeacherID int
	Teacher   User `gorm:"foreignKey:TeacherID"`
	LessonID  int
	Lesson    Lesson `gorm:"foreignKey:LessonID"`
}

func (u *Quiz) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}

func (u *Quiz) BeforeSave(tx *gorm.DB) (err error) {
	if u.ID == 0 {
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()
	u.EndTime = u.StartTime.Add(time.Duration(u.Duration) * time.Minute)
	return
}
