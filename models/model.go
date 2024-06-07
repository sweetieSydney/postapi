package models

import (
	"time"

	"github.com/lib/pq"
)

type Issue struct {
	ID          uint           `gorm:"primaryKey"`
	Title       string         `gorm:"not null"`
	Description string         `gorm:"type:text"`
	Labels      pq.StringArray `gorm:"type:text[]"`
	Assignees   pq.StringArray `gorm:"type:text[]"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
