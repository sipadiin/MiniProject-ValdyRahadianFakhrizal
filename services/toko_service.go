package services

import (
	"errors"

	"evermos/models"
	"evermos/repository"
)

type UpdateTokoInput struct {
	NamaToko string `json:"nama_toko"`
}

func GetAllToko(page, limit int, namaToko string) ([]models.Toko, int64, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	return repository.GetAllToko(limit, offset, namaToko)
}

func GetTokoByID(id uint) (*models.Toko, error) {
	return repository.FindTokoByID(id)
}

func UpdateMyToko(tokoID uint, userID uint, input UpdateTokoInput) (*models.Toko, error) {
	// 1. Cari toko berdasarkan ID
	toko, err := repository.FindTokoByID(tokoID)
	if err != nil {
		return nil, errors.New("toko tidak ditemukan")
	}

	// 2. Validasi: Apakah user yang login adalah pemilik toko ini?
	if toko.UserID != userID {
		return nil, errors.New("anda bukan pemilik toko ini")
	}

	// 3. Update field
	if input.NamaToko != "" {
		toko.NamaToko = input.NamaToko
	}

	if err := repository.UpdateToko(toko); err != nil {
		return nil, errors.New("gagal mengupdate toko")
	}

	return toko, nil
}
