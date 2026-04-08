package repository

import (
	"evermos/config"
	"evermos/models"
)

func GetAllKategori(limit, offset int, nama string) ([]models.Kategori, int64, error) {
	var kategoris []models.Kategori
	var total int64

	db := config.DB.Model(&models.Kategori{})

	if nama != "" {
		db = db.Where("nama_kategori LIKE ?", "%"+nama+"%")
	}

	db.Count(&total)
	err := db.Limit(limit).Offset(offset).Find(&kategoris).Error
	return kategoris, total, err
}

func FindKategoriByID(id uint) (*models.Kategori, error) {
	var kategori models.Kategori
	err := config.DB.First(&kategori, id).Error
	if err != nil {
		return nil, err
	}
	return &kategori, nil
}

func CreateKategori(kategori *models.Kategori) error {
	return config.DB.Create(kategori).Error
}

func UpdateKategori(kategori *models.Kategori) error {
	return config.DB.Save(kategori).Error
}

func DeleteKategori(id uint) error {
	var kategori models.Kategori
	err := config.DB.First(&kategori, id).Error
	if err != nil {
		return err
	}
	return config.DB.Delete(&kategori).Error
}
