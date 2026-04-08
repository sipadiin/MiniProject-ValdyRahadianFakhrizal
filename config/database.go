package config

import (
	"fmt"

	"evermos/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	host := "127.0.0.1"
	port := "3306"
	user := "root"
	password := ""
	dbname := "rakamin"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		panic("Gagal terkoneksi ke database!")
	}

	fmt.Println("Koneksi Database Berhasil!")
	DB = database

	DB.AutoMigrate(
		&models.User{},
		&models.Toko{},
		&models.Alamat{},
		&models.Kategori{},
		&models.Produk{},
		&models.FotoProduk{},
		&models.LogProduk{},
		&models.Trx{},
		&models.DetailTrx{},
	)
	fmt.Println("Auto Migration Berhasil!")
}
