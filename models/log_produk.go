package models

import "time"

type LogProduk struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	IdProduk      uint      `json:"id_produk"`
	IdToko        uint      `json:"id_toko"`
	IdCategory    uint      `json:"id_category"`
	NamaProduk    string    `gorm:"type:varchar(255)" json:"nama_produk"`
	Slug          string    `gorm:"type:varchar(255)" json:"slug"`
	HargaReseller int       `json:"harga_reseller"`
	HargaKonsumen int       `json:"harga_konsumen"`
	Deskripsi     string    `gorm:"type:text" json:"deskripsi"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
