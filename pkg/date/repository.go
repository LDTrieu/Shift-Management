package date

import (
	"errors"

	"gorm.io/gorm"
)

var (
	RecordNotFound = errors.New("record not found")
)

type Repository interface {
	// pass in any date and create the week containing that date if not exist
	Create(date *Date) (*Date, error)
	FindByID(id uint, userID uint) (*Date, error)
	Publish() error
}

type repository struct {
	db *gorm.DB
}

// pass in any date and create the week containing that date if not exist
func (r *repository) Create(date *Date) (*Date, error) {
	err := r.db.Debug().Where("date = ? and user_id = ?", date.Date, date.UserID).First(&date).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		result := r.db.Create(date)
		if result.Error != nil {
			return nil, result.Error
		}
	}

	return date, nil
}

func (r *repository) FindByID(id uint, userID uint) (*Date, error) {
	var date Date
	err := r.db.Debug().Where("id = ? and user_id = ?", id, userID).Preload("Shifts").First(&date).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, RecordNotFound
	}

	return &date, nil
}

func (r *repository) Publish() error {
	panic("not implemented") // TODO: Implement
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}