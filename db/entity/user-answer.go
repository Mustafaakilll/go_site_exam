package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserAnswer struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	UserID   int    `gorm:"user_id"`
	User     User   `gorm:"foreignKey:UserID" json:"user"`
	QuizID   int    `gorm:"quiz_id"`
	Quiz     Quiz   `gorm:"foreignKey:QuizID" json:"quiz"`
	AnswerID int    `gorm:"answer_id"`
	Answer   Answer `gorm:"foreignKey:AnswerID" json:"answer"`
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
