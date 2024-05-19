package entity

import (
	"time"

	"gorm.io/gorm"
)

type CodeSubmission struct {
	ID             int `gorm:"primaryKey"`
	Input          string
	ExpectedOutput string
	CreatedAt      time.Time
	UpdatedAt      time.Time

	CodeID int
	Code   Code `gorm:"foreignKey:CodeID"`
}

func (u *CodeSubmission) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}

func (u *CodeSubmission) BeforeSave(tx *gorm.DB) (err error) {
	if u.ID == 0 {
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()
	return
}
