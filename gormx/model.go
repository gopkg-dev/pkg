package gormx

import (
	"time"

	"gorm.io/gorm"
)

// Model base model
type Model struct {
	ID        int            `gorm:"column:id;primaryKey"`
	CreatedAt time.Time      `gorm:"column:created_at;index;"`
	UpdatedAt time.Time      `gorm:"column:updated_at;index;"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
