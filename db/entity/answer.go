package entity

import (
	"time"

	"gorm.io/gorm"
)

type Answer struct {
	ID        int `gorm:"primaryKey"`
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time

	QuestionID int
	Question   Question `gorm:"foreignKey:QuestionID"`
	ChoiceID   int
	Choice     Choice `gorm:"foreignKey:ChoiceID"`
	UserID     int
	User       User `gorm:"foreignKey:UserID"`
}

func (u *Answer) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}

func (u *Answer) BeforeSave(tx *gorm.DB) (err error) {
	if u.ID == 0 {
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()
	return
}
