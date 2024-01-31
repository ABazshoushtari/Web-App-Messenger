package repository

import (
	"github.com/ABazshoushtari/Web-App-Messenger/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func newUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Register(user *domain.User) error {
	return u.db.Create(&user).Error
}

func (u *userRepository) GetByUsername(username string, user *domain.User) error {
	err := u.db.Where("username = ?", username).First(&user).Error
	return err
}

func (u *userRepository) GetByPhoneNumber(phoneNumber string, user *domain.User) error {
	err := u.db.Where("phone_number = ?", phoneNumber).First(&user).Error
	return err
}

func (u *userRepository) GetByID(id uint64, user *domain.User) error {
	err := u.db.Where("id = ?", id).First(&user).Error
	return err
}

func (u *userRepository) Update(user *domain.User) error {
	return u.db.Save(&user).Error
}

func (u *userRepository) Delete(user *domain.User) error {
	return u.db.Delete(&user).Error
}

func (u *userRepository) GetByKey(key string, value string, user *domain.User) error {
	err := u.db.Where(key+" = ?", value).First(&user).Error
	return err
}
