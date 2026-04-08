package services

import (
	"errors"

	"evermos/models"
	"evermos/repository"
)

type UpdateUserInput struct {
	Nama         string `json:"nama"`
	TanggalLahir string `json:"tanggal_lahir"`
	JenisKelamin string `json:"jenis_kelamin"`
	Tentang      string `json:"tentang"`
	Pekerjaan    string `json:"pekerjaan"`
	IdProvinsi   string `json:"id_provinsi"`
	IdKota       string `json:"id_kota"`
}

func GetMyProfile(userID uint) (*models.User, error) {
	user, err := repository.FindUserByID(userID)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}
	return user, nil
}

func UpdateMyProfile(userID uint, input UpdateUserInput) (*models.User, error) {
	user, err := repository.FindUserByID(userID)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	if input.Nama != "" {
		user.Nama = input.Nama
	}
	if input.TanggalLahir != "" {
		user.TanggalLahir = ParseTanggalLahir(input.TanggalLahir)
	}
	if input.JenisKelamin != "" {
		user.JenisKelamin = input.JenisKelamin
	}
	if input.Tentang != "" {
		user.Tentang = input.Tentang
	}
	if input.Pekerjaan != "" {
		user.Pekerjaan = input.Pekerjaan
	}
	if input.IdProvinsi != "" {
		user.IdProvinsi = input.IdProvinsi
	}
	if input.IdKota != "" {
		user.IdKota = input.IdKota
	}

	if err := repository.UpdateUser(user); err != nil {
		return nil, errors.New("gagal mengupdate profil")
	}

	return user, nil
}
