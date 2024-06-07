package models

import "time"

type Issue struct {
	ID          uint     `gorm:"primaryKey"`
	Title       string   `gorm:"not null"`
	Description string   `gorm:"type:text"`
	Labels      []string `gorm:"type:text[]"`
	Assignees   []string `gorm:"type:text[]"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
