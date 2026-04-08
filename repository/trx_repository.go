package repository

import (
	"evermos/config"
	"evermos/models"

	"gorm.io/gorm"
)

func CreateLogProduk(tx *gorm.DB, log *models.LogProduk) error {
	return tx.Create(log).Error
}

func CreateTrx(tx *gorm.DB, trx *models.Trx) error {
	return tx.Create(trx).Error
}

func CreateDetailTrx(tx *gorm.DB, detail *models.DetailTrx) error {
	return tx.Create(detail).Error
}

func UpdateStokProduk(tx *gorm.DB, produkID uint, kuantitas int) error {
	return tx.Model(&models.Produk{}).Where("id = ?", produkID).
		Update("stok", gorm.Expr("stok - ?", kuantitas)).Error
}

func GetAllTrxByUserID(userID uint, limit, offset int) ([]models.Trx, int64, error) {
	var trxs []models.Trx
	var total int64

	db := config.DB.Model(&models.Trx{}).Where("id_user = ?", userID)
	db.Count(&total)

	err := db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&trxs).Error
	return trxs, total, err
}

func FindTrxByID(id uint) (*models.Trx, error) {
	var trx models.Trx
	err := config.DB.First(&trx, id).Error
	if err != nil {
		return nil, err
	}
	return &trx, nil
}

func GetDetailTrxByTrxID(trxID uint) ([]models.DetailTrx, error) {
	var details []models.DetailTrx
	err := config.DB.Where("id_trx = ?", trxID).Find(&details).Error
	return details, err
}

func FindLogProdukByID(id uint) (*models.LogProduk, error) {
	var log models.LogProduk
	err := config.DB.First(&log, id).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}

func CountTrxToday() int64 {
	var count int64
	config.DB.Model(&models.Trx{}).
		Where("DATE(created_at) = CURDATE()").
		Count(&count)
	return count
}

func BeginTransaction() *gorm.DB {
	return config.DB.Begin()
}
