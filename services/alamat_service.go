package services

import (
	"errors"

	"evermos/models"
	"evermos/repository"
)

func GetAllMyAlamat(userID uint, limit, offset int) ([]models.Alamat, int64, error) {
	return repository.GetAlamatByUserID(userID, limit, offset)
}

func CreateNewAlamat(userID uint, input models.Alamat) (models.Alamat, error) {
	input.IdUser = userID
	err := repository.CreateAlamat(&input)
	return input, err
}

func UpdateMyAlamat(id, userID uint, input models.Alamat) (*models.Alamat, error) {
	alamat, err := repository.FindAlamatByID(id)
	if err != nil {
		return nil, errors.New("alamat tidak ditemukan")
	}
	if alamat.IdUser != userID {
		return nil, errors.New("alamat bukan milik anda")
	}

	if input.JudulAlamat != "" {
		alamat.JudulAlamat = input.JudulAlamat
	}
	if input.NamaPenerima != "" {
		alamat.NamaPenerima = input.NamaPenerima
	}
	if input.NoTelp != "" {
		alamat.NoTelp = input.NoTelp
	}
	if input.DetailAlamat != "" {
		alamat.DetailAlamat = input.DetailAlamat
	}

	err = repository.UpdateAlamat(alamat)
	return alamat, err
}

func DeleteMyAlamat(id, userID uint) error {
	alamat, err := repository.FindAlamatByID(id)
	if err != nil {
		return errors.New("alamat tidak ditemukan")
	}
	if alamat.IdUser != userID {
		return errors.New("alamat bukan milik anda")
	}
	return repository.DeleteAlamat(id)
}
