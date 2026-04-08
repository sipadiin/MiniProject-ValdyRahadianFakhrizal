package services

import (
	"errors"

	"evermos/models"
	"evermos/repository"

	"golang.org/x/crypto/bcrypt"
)

// RegisterInput adalah struct untuk menerima data register dari controller
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

// LoginInput adalah struct untuk menerima data login dari controller
type LoginInput struct {
	Notelp    string `json:"no_telp" validate:"required"`
	Katasandi string `json:"kata_sandi" validate:"required"`
}

// Register membuat user baru + otomatis buat toko
func Register(input RegisterInput) (*models.User, error) {
	// 1. Cek apakah email sudah terdaftar
	existingUser, err := repository.FindUserByEmail(input.Email)
	if err != nil {
		return nil, errors.New("gagal mengecek email")
	}
	if existingUser != nil {
		return nil, errors.New("email sudah terdaftar")
	}

	// 2. Cek apakah notelp sudah terdaftar
	existingUser, err = repository.FindUserByNoTelp(input.Notelp)
	if err != nil {
		return nil, errors.New("gagal mengecek nomor telepon")
	}
	if existingUser != nil {
		return nil, errors.New("nomor telepon sudah terdaftar")
	}

	// 3. Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Katasandi), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("gagal mengenkripsi password")
	}

	// 4. Buat user baru
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

	// 5. Simpan user ke database
	if err := repository.CreateUser(&user); err != nil {
		return nil, errors.New("gagal membuat user")
	}

	// 6. Otomatis buat toko untuk user
	toko := models.Toko{
		UserID:   user.ID,
		NamaToko: "Toko " + user.Nama,
	}

	if err := repository.CreateToko(&toko); err != nil {
		return nil, errors.New("gagal membuat toko")
	}

	// 7. Set toko ke user untuk response
	user.Toko = toko

	return &user, nil
}

// Login memverifikasi kredensial dan mengembalikan user
func Login(input LoginInput) (*models.User, error) {
	// 1. Cari user berdasarkan notelp (sudah include Preload Toko)
	user, err := repository.FindUserByNoTelp(input.Notelp)
	if err != nil {
		return nil, errors.New("terjadi kesalahan saat mencari user")
	}
	if user == nil {
		return nil, errors.New("nomor telepon tidak ditemukan")
	}

	// 2. Bandingkan password
	err = bcrypt.CompareHashAndPassword([]byte(user.Katasandi), []byte(input.Katasandi))
	if err != nil {
		return nil, errors.New("kata sandi salah")
	}

	return user, nil
}
