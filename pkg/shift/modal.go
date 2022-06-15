package shift

import (
	"time"

	"gorm.io/gorm"
)

type Shift struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	DateID uint `json:"date_id"`
	UserID uint `json:"user_id"`

	QuarterStart uint `json:"quarter_start"`
	NumQuarter   uint `json:"num_quarter"`

	Title       string `json:"title"`
	Description string `json:"description"`
}