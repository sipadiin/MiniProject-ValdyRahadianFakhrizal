package services

import (
	"evermos/models"
	"evermos/repository"
)

func ListKategori(limit, offset int, nama string) ([]models.Kategori, int64, error) {
	return repository.GetAllKategori(limit, offset, nama)
}

func DetailKategori(id uint) (*models.Kategori, error) {
	return repository.FindKategoriByID(id)
}

func AddKategori(nama string) (models.Kategori, error) {
	kat := models.Kategori{NamaKategori: nama}
	err := repository.CreateKategori(&kat)
	return kat, err
}

func EditKategori(id uint, nama string) (*models.Kategori, error) {
	kat, err := repository.FindKategoriByID(id)
	if err != nil {
		return nil, err
	}

	kat.NamaKategori = nama
	err = repository.UpdateKategori(kat)
	return kat, err
}

func RemoveKategori(id uint) error {
	return repository.DeleteKategori(id)
}
