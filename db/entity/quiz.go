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
	DeletedAt   gorm.DeletedAt

	TeacherID int    `json:"teacher_id"`
	Teacher   User   `gorm:"foreignKey:TeacherID" json:"teacher"`
	LessonID  int    `json:"lesson_id"`
	Lesson    Lesson `gorm:"foreignKey:LessonID" json:"lesson"`
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
