package entity

import (
	"time"

	"gorm.io/gorm"
)

type CodeAnswer struct {
	ID        int `gorm:"primaryKey"`
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	UserID int  `json:"user_id"`
	User   User `gorm:"foreignKey:UserID" json:"user"`
	CodeID int  `json:"code_id"`
	Code   Code `gorm:"foreignKey:CodeID" json:"code"`
}

func (u *CodeAnswer) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}

func (u *CodeAnswer) BeforeSave(tx *gorm.DB) (err error) {
	if u.ID == 0 {
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()
	return
}
