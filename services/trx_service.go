package services

import (
	"errors"
	"fmt"
	"time"

	"evermos/models"
	"evermos/repository"
)

// TrxItemInput mewakili satu produk dalam transaksi
type TrxItemInput struct {
	ProductID uint `json:"product_id"` // dari id_produk ke product_id
	Kuantitas int  `json:"kuantitas"`
}

// CreateTrxInput mewakili input untuk membuat transaksi baru
type CreateTrxInput struct {
	AlamatKirim uint           `json:"alamat_kirim"` // dari alamat_pengiriman ke alamat_kirim
	MethodBayar string         `json:"method_bayar"`
	DetailTrx   []TrxItemInput `json:"detail_trx"`
}

// TrxDetailResponse untuk response detail transaksi
type TrxDetailResponse struct {
	Trx       models.Trx      `json:"trx"`
	DetailTrx []DetailTrxItem `json:"detail_trx"`
}

// DetailTrxItem untuk response per item di detail transaksi
type DetailTrxItem struct {
	models.DetailTrx
	LogProduk models.LogProduk `json:"log_produk"`
}

// CreateTransaksi membuat transaksi baru dengan DB Transaction
func CreateTransaksi(userID uint, input CreateTrxInput) (*TrxDetailResponse, error) {
	// 1. Validasi list produk tidak kosong
	if len(input.DetailTrx) == 0 {
		return nil, errors.New("detail transaksi tidak boleh kosong")
	}

	// 2. Validasi alamat milik user sendiri
	alamat, err := repository.FindAlamatByID(input.AlamatKirim)
	if err != nil {
		return nil, errors.New("alamat tidak ditemukan")
	}
	if alamat.IdUser != userID {
		return nil, errors.New("alamat bukan milik anda")
	}

	// 3. Mulai DB Transaction
	tx := repository.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var totalHarga int
	var detailResponses []DetailTrxItem

	// 4. Generate kode invoice: INV-YYYYMMDD-XXX
	today := time.Now().Format("20060102")
	count := repository.CountTrxToday()
	kodeInvoice := fmt.Sprintf("INV-%s-%03d", today, count+1)

	// 5. Buat transaksi dulu (harga_total diupdate nanti)
	trx := models.Trx{
		IdUser:           userID,
		AlamatPengiriman: input.AlamatKirim,
		HargaTotal:       0,
		KodeInvoice:      kodeInvoice,
		MethodBayar:      input.MethodBayar,
	}

	if err := repository.CreateTrx(tx, &trx); err != nil {
		tx.Rollback()
		return nil, errors.New("gagal membuat transaksi")
	}

	// 6. Proses setiap produk
	for _, item := range input.DetailTrx {
		// a. Cari produk
		produk, err := repository.FindProdukByID(item.ProductID)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("produk dengan ID %d tidak ditemukan", item.ProductID)
		}

		// b. Cek stok cukup
		if produk.Stok < item.Kuantitas {
			tx.Rollback()
			return nil, fmt.Errorf("stok produk '%s' tidak cukup (sisa: %d, diminta: %d)",
				produk.NamaProduk, produk.Stok, item.Kuantitas)
		}

		// c. Buat log_produk (snapshot data produk saat ini)
		logProduk := models.LogProduk{
			IdProduk:      produk.ID,
			IdToko:        produk.IdToko,
			IdCategory:    produk.IdCategory,
			NamaProduk:    produk.NamaProduk,
			Slug:          produk.Slug,
			HargaReseller: produk.HargaReseller,
			HargaKonsumen: produk.HargaKonsumen,
			Deskripsi:     produk.Deskripsi,
		}

		if err := repository.CreateLogProduk(tx, &logProduk); err != nil {
			tx.Rollback()
			return nil, errors.New("gagal membuat log produk")
		}

		// d. Kurangi stok produk
		if err := repository.UpdateStokProduk(tx, produk.ID, item.Kuantitas); err != nil {
			tx.Rollback()
			return nil, errors.New("gagal mengurangi stok produk")
		}

		// e. Hitung harga total per item
		hargaItem := produk.HargaKonsumen * item.Kuantitas
		totalHarga += hargaItem

		// f. Buat detail transaksi
		detail := models.DetailTrx{
			IdTrx:       trx.Id,
			IdLogProduk: logProduk.ID,
			IdToko:      produk.IdToko,
			Kuantitas:   item.Kuantitas,
			HargaTotal:  hargaItem,
		}

		if err := repository.CreateDetailTrx(tx, &detail); err != nil {
			tx.Rollback()
			return nil, errors.New("gagal membuat detail transaksi")
		}

		detailResponses = append(detailResponses, DetailTrxItem{
			DetailTrx: detail,
			LogProduk: logProduk,
		})
	}

	// 7. Update total harga di transaksi
	trx.HargaTotal = totalHarga
	if err := tx.Model(&trx).Update("harga_total", totalHarga).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("gagal mengupdate total harga")
	}

	// 8. Commit transaction
	if err := tx.Commit().Error; err != nil {
		return nil, errors.New("gagal menyimpan transaksi")
	}

	return &TrxDetailResponse{
		Trx:       trx,
		DetailTrx: detailResponses,
	}, nil
}

// GetAllTransaksi mengambil semua transaksi milik user
func GetAllTransaksi(userID uint, limit, offset int) ([]models.Trx, int64, error) {
	return repository.GetAllTrxByUserID(userID, limit, offset)
}

// GetDetailTransaksi mengambil detail transaksi + log produk
func GetDetailTransaksi(userID uint, trxID uint) (*TrxDetailResponse, error) {
	// 1. Cari transaksi
	trx, err := repository.FindTrxByID(trxID)
	if err != nil {
		return nil, errors.New("transaksi tidak ditemukan")
	}

	// 2. Validasi kepemilikan
	if trx.IdUser != userID {
		return nil, errors.New("transaksi bukan milik anda")
	}

	// 3. Ambil detail transaksi
	details, err := repository.GetDetailTrxByTrxID(trxID)
	if err != nil {
		return nil, errors.New("gagal mengambil detail transaksi")
	}

	// 4. Ambil log produk untuk setiap detail
	var detailItems []DetailTrxItem
	for _, d := range details {
		logProduk, _ := repository.FindLogProdukByID(d.IdLogProduk)
		item := DetailTrxItem{
			DetailTrx: d,
		}
		if logProduk != nil {
			item.LogProduk = *logProduk
		}
		detailItems = append(detailItems, item)
	}

	return &TrxDetailResponse{
		Trx:       *trx,
		DetailTrx: detailItems,
	}, nil
}
