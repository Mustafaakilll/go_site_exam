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
	DeletedAt gorm.DeletedAt

	QuestionID int      `json:"question_id"`
	Question   Question `gorm:"foreignKey:QuestionID" json:"question"`
	ChoiceID   int      `json:"choice_id"`
	Choice     Choice   `gorm:"foreignKey:ChoiceID" json:"choice"`
	UserID     int      `json:"user_id"`
	User       User     `gorm:"foreignKey:UserID" json:"user"`
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
