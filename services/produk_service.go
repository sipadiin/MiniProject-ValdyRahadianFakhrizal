package services

import (
	"errors"
	"strings"

	"evermos/models"
	"evermos/repository"
)

func CreateNewProduk(userID uint, input models.Produk, fotoNames []string) (models.Produk, error) {
	toko, err := repository.FindTokoByID(userID)
	if err != nil {
		return models.Produk{}, err
	}

	_, err = repository.FindKategoriByID(input.IdCategory)
	if err != nil {
		return models.Produk{}, errors.New("kategori tidak ditemukan")
	}

	input.IdToko = toko.ID
	input.Slug = strings.ToLower(strings.ReplaceAll(input.NamaProduk, " ", "-"))

	if err := repository.CreateProduk(&input); err != nil {
		return models.Produk{}, err
	}

	for _, name := range fotoNames {
		foto := models.FotoProduk{
			IdProduk: input.ID,
			Url:      "uploads/" + name,
		}
		repository.CreateFotoProduk(&foto)
	}

	return input, nil
}
