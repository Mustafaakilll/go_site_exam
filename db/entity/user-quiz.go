package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserQuiz struct {
	ID        int `gorm:"primaryKey"`
	Result    int
	IsReview  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	UserID int  `json:"user_id"`
	User   User `gorm:"foreignKey:UserID" json:"user"`
	QuizID int  `json:"quiz_id"`
	Quiz   Quiz `gorm:"foreignKey:QuizID" json:"quiz"`
}

func (u *UserQuiz) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}

func (u *UserQuiz) BeforeSave(tx *gorm.DB) (err error) {
	if u.ID == 0 {
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()
	return
}
