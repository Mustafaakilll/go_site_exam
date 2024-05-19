package entity

import (
	"time"

	"gorm.io/gorm"
)

type Choice struct {
	ID        int `gorm:"primaryKey"`
	Text      string
	IsCorrect bool
	CreatedAt time.Time
	UpdatedAt time.Time

	QuestionID int
	Question   Question `gorm:"foreignKey:QuestionID"`
}

func (u *Choice) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}

func (u *Choice) BeforeSave(tx *gorm.DB) (err error) {
	if u.ID == 0 {
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()
	return
}
