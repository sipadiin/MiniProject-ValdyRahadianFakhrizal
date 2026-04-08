package models

import "time"

type DetailTrx struct {
	Id          uint      `gorm:"primaryKey" json:"id"`
	IdTrx       uint      `json:"id_trx"`
	IdLogProduk uint      `json:"id_log_produk"`
	IdToko      uint      `json:"id_toko"`
	Kuantitas   int       `json:"kuantitas"`
	HargaTotal  int       `json:"harga_total"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}
