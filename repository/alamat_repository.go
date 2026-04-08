package repository

import (
	"evermos/config"
	"evermos/models"
)

func CreateAlamat(alamat *models.Alamat) error {
	return config.DB.Create(alamat).Error
}

func GetAlamatByUserID(userID uint, limit, offset int) ([]models.Alamat, int64, error) {
	var alamats []models.Alamat
	var total int64

	db := config.DB.Model(&models.Alamat{}).Where("id_user = ?", userID)
	db.Count(&total)

	err := db.Limit(limit).Offset(offset).Find(&alamats).Error
	return alamats, total, err
}

func FindAlamatByID(id uint) (*models.Alamat, error) {
	var alamat models.Alamat
	err := config.DB.First(&alamat, id).Error
	if err != nil {
		return nil, err
	}
	return &alamat, nil
}

func UpdateAlamat(alamat *models.Alamat) error {
	return config.DB.Save(alamat).Error
}

func DeleteAlamat(id uint) error {
	var alamat models.Alamat
	err := config.DB.First(&alamat, id).Error
	if err != nil {
		return err
	}
	return config.DB.Delete(&alamat).Error
}
