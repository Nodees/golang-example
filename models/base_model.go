package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	CreatedAt time.Time
	UpdatedAt time.Time
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
