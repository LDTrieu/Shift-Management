package week

import(
	"errors"
	"time"

	"gorm.io/gorm"
	"github.com/ldtrieu/staffany-backend/pkg/utils"
)

var (
	RecordNotFound = errors.New("record not found")
)

type Repository interface {
	// pass in any date and create the week containing that date if not exist
	Create(date string, userID uint) (id uint, err error)
	// input a date and return the week containing that date
	FindByDateAndUserID(date string, userID uint) (*Week, error)
	Publish() error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// date string => find week range
func (r *repository) Create(date string, userID uint) (id uint, err error) {
	weekNumber, startDate := utils.GetWeekFromDateString(date)
	week := Week{
		WeekNumber: weekNumber,
		UserID:     userID,
		StartDate:  utils.DateStringToInt(startDate),
	}

	result := r.db.Create(&week)
	if result.Error != nil {
		return 0, result.Error
	}

	return week.ID, nil
}

func (r *repository) GetCurrentWeek(userID uint) (*Week, error) {
	now := time.Now()
	date := utils.GetDateString(now)
	week, err := r.FindByDateAndUserID(date, userID)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, RecordNotFound
	}

	return week, nil
}

// input a date and return the week containing that date
func (r *repository) FindByDateAndUserID(date string, userID uint) (*Week, error) {
	_, startDate := utils.GetWeekFromDateString(date)
	dateInt := utils.DateStringToInt(startDate)

	var week Week
	result := r.db.Debug().Where("start_date = ? AND user_id = ?", dateInt, userID).Preload("Date.Shifts").First(&week)
	err := result.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, RecordNotFound
		}

		return nil, err

	}

	return &week, nil
}

func (r *repository) Publish() error {
	panic("not implemented")
}