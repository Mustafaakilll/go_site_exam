package entity

import (
	"time"

	"gorm.io/gorm"
)

type UserType struct {
	ID   int `gorm:"primaryKey"`
	Name string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *UserType) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}

func (u *UserType) BeforeSave(tx *gorm.DB) (err error) {
	if u.ID == 0 {
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()
	return
}
