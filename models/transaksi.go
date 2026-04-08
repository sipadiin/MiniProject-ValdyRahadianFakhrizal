package models

import "time"

type Trx struct {
	Id               uint      `gorm:"primaryKey" json:"id"`
	IdUser           uint      `json:"id_user"`
	AlamatPengiriman uint      `json:"alamat_pengiriman"`
	HargaTotal       int       `json:"harga_total"`
	KodeInvoice      string    `gorm:"type:varchar(255)" json:"kode_invoice"`
	MethodBayar      string    `gorm:"type:varchar(255)" json:"method_bayar"`
	UpdatedAt        time.Time `json:"updated_at"`
	CreatedAt        time.Time `json:"created_at"`
}
