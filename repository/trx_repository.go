package repository

import (
	"evermos/config"
	"evermos/models"

	"gorm.io/gorm"
)

// CreateLogProduk menyimpan snapshot produk saat transaksi
func CreateLogProduk(tx *gorm.DB, log *models.LogProduk) error {
	return tx.Create(log).Error
}

// CreateTrx menyimpan transaksi baru
func CreateTrx(tx *gorm.DB, trx *models.Trx) error {
	return tx.Create(trx).Error
}

// CreateDetailTrx menyimpan detail transaksi
func CreateDetailTrx(tx *gorm.DB, detail *models.DetailTrx) error {
	return tx.Create(detail).Error
}

// UpdateStokProduk mengurangi stok produk
func UpdateStokProduk(tx *gorm.DB, produkID uint, kuantitas int) error {
	return tx.Model(&models.Produk{}).Where("id = ?", produkID).
		Update("stok", gorm.Expr("stok - ?", kuantitas)).Error
}

// GetAllTrxByUserID mengambil semua transaksi milik user dengan pagination
func GetAllTrxByUserID(userID uint, limit, offset int) ([]models.Trx, int64, error) {
	var trxs []models.Trx
	var total int64

	db := config.DB.Model(&models.Trx{}).Where("id_user = ?", userID)
	db.Count(&total)

	err := db.Limit(limit).Offset(offset).Order("created_at DESC").Find(&trxs).Error
	return trxs, total, err
}

// FindTrxByID mencari transaksi berdasarkan ID
func FindTrxByID(id uint) (*models.Trx, error) {
	var trx models.Trx
	err := config.DB.First(&trx, id).Error
	if err != nil {
		return nil, err
	}
	return &trx, nil
}

// GetDetailTrxByTrxID mengambil semua detail transaksi berdasarkan trx ID
func GetDetailTrxByTrxID(trxID uint) ([]models.DetailTrx, error) {
	var details []models.DetailTrx
	err := config.DB.Where("id_trx = ?", trxID).Find(&details).Error
	return details, err
}

// FindLogProdukByID mencari log produk berdasarkan ID
func FindLogProdukByID(id uint) (*models.LogProduk, error) {
	var log models.LogProduk
	err := config.DB.First(&log, id).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}

// CountTrxToday menghitung jumlah transaksi hari ini (untuk generate invoice)
func CountTrxToday() int64 {
	var count int64
	config.DB.Model(&models.Trx{}).
		Where("DATE(created_at) = CURDATE()").
		Count(&count)
	return count
}

// BeginTransaction memulai database transaction
func BeginTransaction() *gorm.DB {
	return config.DB.Begin()
}
