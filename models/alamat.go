package models

import "time"

type Alamat struct {
	Id           uint      `gorm:"primaryKey" json:"id"`
	IdUser       uint      `json:"id_user"` // FK ke User
	JudulAlamat  string    `gorm:"type:varchar(255)" json:"judul_alamat"`
	NamaPenerima string    `gorm:"type:varchar(255)" json:"nama_penerima"`
	NoTelp       string    `gorm:"type:varchar(255)" json:"no_telp"`
	DetailAlamat string    `gorm:"type:text" json:"detail_alamat"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
}
