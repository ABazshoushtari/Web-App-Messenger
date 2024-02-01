package repository

import (
	"github.com/ABazshoushtari/Web-App-Messenger/domain"
	"github.com/ABazshoushtari/Web-App-Messenger/internal/logger"
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
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

func (u *userRepository) SetImage(user *domain.User, image *multipart.FileHeader) error {
	imgNameSplitted := strings.Split(image.Filename, ".")
	imgFormat := imgNameSplitted[len(imgNameSplitted)-1]
	image.Filename = strconv.FormatUint(user.ID, 10) + "." + imgFormat
	if _, err := os.Stat("./assets"); os.IsNotExist(err) {
		err := os.MkdirAll("./assets", 0755)
		if err != nil {
			logger.Logger().Errorw("Error while Creating directory to save user image", "error", err)
			return err
		}
	}
	savedFile, err := os.Create("./assets/" + image.Filename)

	if err != nil {
		logger.Logger().Errorw("Error while creating image for saving it into assets", "error", err)
		return err
	}
	defer savedFile.Close()

	imgFile, err := image.Open()
	defer imgFile.Close()
	if err != nil {
		logger.Logger().Errorw("Error while Opening image for saving it into assets", "error", err)
		return err
	}

	imgData, err := io.ReadAll(imgFile)
	if err != nil {
		logger.Logger().Errorw("Error while Reading image for saving it into assets", "error", err)
		return err
	}
	if _, err := savedFile.Write(imgData); err != nil {
		logger.Logger().Errorw("Error while saving image for saving it into assets", "error", err)
		return err
	}

	return u.db.Model(&user).Update("image", image.Filename).Error
}
func (u *userRepository) Delete(user *domain.User) error {
	return u.db.Delete(&user).Error
}

func (u *userRepository) GetByKey(value string, user *domain.User) error {
	err := u.db.Where("username = ? or phone_number = ?", value).First(&user).Error
	return err
}

func (u *userRepository) CheckExisting(username string, phoneNumber string) error {
	return u.db.Where("username = ? or phone_number = ?", username, phoneNumber).First(&domain.User{}).Error
}
