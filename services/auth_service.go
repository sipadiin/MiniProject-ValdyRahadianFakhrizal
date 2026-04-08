package services

import (
	"errors"

	"evermos/models"
	"evermos/repository"

	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Nama         string `json:"nama" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Notelp       string `json:"no_telp" validate:"required"`
	Katasandi    string `json:"kata_sandi" validate:"required,min=6"`
	TanggalLahir string `json:"tanggal_lahir"`
	JenisKelamin string `json:"jenis_kelamin"`
	Tentang      string `json:"tentang"`
	Pekerjaan    string `json:"pekerjaan"`
	IdProvinsi   string `json:"id_provinsi"`
	IdKota       string `json:"id_kota"`
}

type LoginInput struct {
	Notelp    string `json:"no_telp" validate:"required"`
	Katasandi string `json:"kata_sandi" validate:"required"`
}

func Register(input RegisterInput) (*models.User, error) {
	existingUser, err := repository.FindUserByEmail(input.Email)
	if err != nil {
		return nil, errors.New("gagal mengecek email")
	}
	if existingUser != nil {
		return nil, errors.New("email sudah terdaftar")
	}

	existingUser, err = repository.FindUserByNoTelp(input.Notelp)
	if err != nil {
		return nil, errors.New("gagal mengecek nomor telepon")
	}
	if existingUser != nil {
		return nil, errors.New("nomor telepon sudah terdaftar")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Katasandi), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("gagal mengenkripsi password")
	}

	user := models.User{
		Nama:         input.Nama,
		Email:        input.Email,
		Notelp:       input.Notelp,
		Katasandi:    string(hashedPassword),
		TanggalLahir: ParseTanggalLahir(input.TanggalLahir),
		JenisKelamin: input.JenisKelamin,
		Tentang:      input.Tentang,
		Pekerjaan:    input.Pekerjaan,
		IdProvinsi:   input.IdProvinsi,
		IdKota:       input.IdKota,
		IsAdmin:      false,
	}

	if err := repository.CreateUser(&user); err != nil {
		return nil, errors.New("gagal membuat user")
	}

	toko := models.Toko{
		UserID:   user.ID,
		NamaToko: "Toko " + user.Nama,
	}

	if err := repository.CreateToko(&toko); err != nil {
		return nil, errors.New("gagal membuat toko")
	}

	user.Toko = toko

	return &user, nil
}

func Login(input LoginInput) (*models.User, error) {
	user, err := repository.FindUserByNoTelp(input.Notelp)
	if err != nil {
		return nil, errors.New("terjadi kesalahan saat mencari user")
	}
	if user == nil {
		return nil, errors.New("nomor telepon tidak ditemukan")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Katasandi), []byte(input.Katasandi))
	if err != nil {
		return nil, errors.New("kata sandi salah")
	}

	return user, nil
}
