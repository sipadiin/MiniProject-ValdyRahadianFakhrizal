package repository

import (
	"errors"

	"evermos/config"
	"evermos/models"

	"gorm.io/gorm"
)

// CreateUser menyimpan user baru ke database
func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

// FindUserByEmail mencari user berdasarkan email
// Return nil, nil jika tidak ditemukan (bukan error)
func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Tidak ditemukan = bukan error
		}
		return nil, err
	}
	return &user, nil
}

// FindUserByNoTelp mencari user berdasarkan nomor telepon
// Return nil, nil jika tidak ditemukan (bukan error)
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

// FindUserByID mencari user berdasarkan ID (dengan preload Toko)
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

// UpdateUser mengupdate data user
func UpdateUser(user *models.User) error {
	return config.DB.Save(user).Error
}
