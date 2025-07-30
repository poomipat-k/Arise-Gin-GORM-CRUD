package models

import (
	"time"

	"gorm.io/gorm"
)

// Item represents a model in the database

type Item struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name" example:"Sample Item"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty"`
}
