package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:dt_created_at"`
	UpdatedAt time.Time `gorm:"column:dt_updated_at"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	base.CreatedAt = time.Now()
	base.UpdatedAt = time.Now()
	return nil
}

func (base *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	base.UpdatedAt = time.Now()
	return nil
}
