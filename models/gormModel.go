package models

import (
	"time"

	"gorm.io/gorm"
)

type GormModel struct {
	ID        uint64         `gorm:"primaryKey" json:"id"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}
