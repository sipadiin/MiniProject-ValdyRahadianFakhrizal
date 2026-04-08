package repository

import (
	"evermos/config"
	"evermos/models"
)

func GetAllToko(limit, offset int, namaToko string) ([]models.Toko, int64, error) {
	var tokos []models.Toko
	var total int64

	db := config.DB.Model(&models.Toko{})

	if namaToko != "" {
		db = db.Where("nama_toko LIKE ?", "%"+namaToko+"%")
	}

	db.Count(&total)

	err := db.Limit(limit).Offset(offset).Find(&tokos).Error
	return tokos, total, err
}

func FindTokoByID(id uint) (*models.Toko, error) {
	var toko models.Toko
	err := config.DB.First(&toko, id).Error
	if err != nil {
		return nil, err
	}
	return &toko, nil
}

func UpdateToko(toko *models.Toko) error {
	return config.DB.Save(toko).Error
}

func CreateToko(toko *models.Toko) error {
	return config.DB.Create(toko).Error
}

func FindTokoByUserID(userID uint) (*models.Toko, error) {
	var toko models.Toko
	err := config.DB.Where("user_id = ?", userID).First(&toko).Error
	if err != nil {
		return nil, err
	}
	return &toko, nil
}
