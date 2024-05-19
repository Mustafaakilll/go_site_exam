package entity

import (
	"time"

	"gorm.io/gorm"
)

type Question struct {
	ID        int `gorm:"primaryKey"`
	Text      string
	Type      int
	Point     int
	CreatedAt time.Time
	UpdatedAt time.Time

	QuizID int
	Quiz   Quiz `gorm:"foreignKey:QuizID"`
}

func (u *Question) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}

func (u *Question) BeforeSave(tx *gorm.DB) (err error) {
	if u.ID == 0 {
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()
	return
}
