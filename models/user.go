package models

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Nama         string    `gorm:"type:varchar(255)" json:"nama"`
	Email        string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	Notelp       string    `gorm:"column:no_telp;type:varchar(20);unique;not null" json:"no_telp"`
	Katasandi    string    `gorm:"type:varchar(255);not null" json:"-"`
	IsAdmin      bool      `gorm:"type:boolean" json:"is_admin"`
	TanggalLahir time.Time `gorm:"type:date" json:"tanggal_lahir"`
	JenisKelamin string    `gorm:"type:varchar(20)" json:"jenis_kelamin"`
	Tentang      string    `gorm:"type:text" json:"tentang"`
	Pekerjaan    string    `gorm:"type:varchar(255)" json:"pekerjaan"`
	IdProvinsi   string    `gorm:"type:varchar(10)" json:"id_provinsi"`
	IdKota       string    `gorm:"type:varchar(10)" json:"id_kota"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Toko Toko `gorm:"foreignKey:UserID" json:"toko"`
}

type Toko struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	NamaToko  string    `gorm:"type:varchar(255)" json:"nama_toko"`
	UrlFoto   string    `gorm:"type:varchar(255)" json:"url_foto"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
