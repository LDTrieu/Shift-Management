package user

import (
	"errors"

	"gorm.io/gorm"
)

var (
	RecordNotFound = errors.New("record not found")
)

type Repository interface {
	// pass in any date and create the week containing that date if not exist
	Create(username, password string) (id uint, err error)
	FindByUserName(username string) (*User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(username, password string) (id uint, err error) {
	user := User{
		UserName: username,
		// should hash password here
		Password: password,
	}

	result := r.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}

func (r *repository) FindByUserName(username string) (*User, error) {
	var user User
	result := r.db.Where("user_name = ?", username).First(&user)
	err := result.Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, RecordNotFound
		}

		return nil, err
	}

	return &user, nil

}