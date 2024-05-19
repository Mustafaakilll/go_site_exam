package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserAnswer struct {
	ID       int `gorm:"primaryKey"`
	UserID   int
	User     User `gorm:"foreignKey:UserID"`
	QuizID   int
	Quiz     Quiz `gorm:"foreignKey:QuizID"`
	AnswerID int
	Answer   Answer `gorm:"foreignKey:AnswerID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *UserAnswer) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}

func (u *UserAnswer) BeforeSave(tx *gorm.DB) (err error) {
	if u.ID == 0 {
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()
	return
}
