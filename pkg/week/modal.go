package week

import(
	"time"

	"github.com/ldtrieu/staffany-backend/pkg/date"
	"gorm.io/gorm"

)

type Week struct{
	ID        uint           	`gorm:"primarykey" json:"id"`
	CreatedAt time.Time     	`json:"created_at"`
	UpdatedAt time.Time      	`json:"updated_at"`
	DeletedAt gorm.DeletedAt	`gorm:"index" json:"-"`

	WeekNumber int 				`json:"week_number"`
	StartDate uint64 			`json:"start_date"`

	IsPublished bool 			`json:"is_published"`
	UserID      uint 			`json:"user_id"`

	Date		[]date.Date		`json:"dates"`
}