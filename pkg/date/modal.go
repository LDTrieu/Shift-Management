package date

import (
	"time"

	"gorm.io/gorm"
	"github.com/ldtrieu/staffany-backend/pkg/shift"
)

type Date struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Date uint64 `json:"date"`

	WeekID uint `json:"week_id"`
	UserID uint `json:"user_id"`

	IsPublished bool `json:"is_published"`

	Shifts []shift.Shift `json:"shifts"`
}