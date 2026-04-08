package models

import "time"

type Kategori struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	NamaKategori string    `gorm:"type:varchar(255)" json:"nama_kategori"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
