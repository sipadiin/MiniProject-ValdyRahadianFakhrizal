package repository

import (
	"evermos/config"
	"evermos/models"
)

func CreateProduk(produk *models.Produk) error {
	return config.DB.Create(produk).Error
}

func CreateFotoProduk(foto *models.FotoProduk) error {
	return config.DB.Create(foto).Error
}

func GetAllProduk(limit, offset int, filters map[string]interface{}) ([]models.Produk, int64, error) {
	var produks []models.Produk
	var total int64

	db := config.DB.Model(&models.Produk{}).Preload("FotoProduk")

	if n, ok := filters["nama"]; ok && n != "" {
		db = db.Where("nama_produk LIKE ?", "%"+n.(string)+"%")
	}
	if c, ok := filters["id_category"]; ok && c != "" {
		db = db.Where("id_category = ?", c)
	}
	if t, ok := filters["id_toko"]; ok && t != "" {
		db = db.Where("id_toko = ?", t)
	}
	if min, ok := filters["min_harga"]; ok && min != "" {
		db = db.Where("harga_konsumen >= ?", min)
	}
	if max, ok := filters["max_harga"]; ok && max != "" {
		db = db.Where("harga_konsumen <= ?", max)
	}

	db.Count(&total)
	err := db.Limit(limit).Offset(offset).Find(&produks).Error
	return produks, total, err
}

func FindProdukByID(id uint) (*models.Produk, error) {
	var produk models.Produk
	err := config.DB.Preload("FotoProduk").First(&produk, id).Error
	return &produk, err
}

func UpdateProduk(produk *models.Produk) error {
	return config.DB.Save(produk).Error
}

func DeleteProduk(id uint) error {
	config.DB.Where("id_produk = ?", id).Delete(&models.FotoProduk{})
	return config.DB.Delete(&models.Produk{}, id).Error
}
