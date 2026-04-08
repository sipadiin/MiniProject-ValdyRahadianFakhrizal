package repository

import (
	"errors"

	"evermos/config"
	"evermos/models"

	"gorm.io/gorm"
)

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func FindUserByNoTelp(notelp string) (*models.User, error) {
	var user models.User
	err := config.DB.Preload("Toko").Where("no_telp = ?", notelp).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func FindUserByID(id uint) (*models.User, error) {
	var user models.User
	err := config.DB.Preload("Toko").First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func UpdateUser(user *models.User) error {
	return config.DB.Save(user).Error
}
