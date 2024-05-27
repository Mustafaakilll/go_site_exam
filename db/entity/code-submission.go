package entity

import (
	"time"

	"gorm.io/gorm"
)

type CodeSubmission struct {
	ID             int `gorm:"primaryKey"`
	Input          string
	ExpectedOutput string `json:"expected_output"`
	CreatedAt      time.Time
	UpdatedAt      time.Time

	CodeID int  `json:"code_id"`
	Code   Code `gorm:"foreignKey:CodeID" json:"code"`
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
