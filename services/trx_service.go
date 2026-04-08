package services

import (
	"errors"
	"fmt"
	"time"

	"evermos/models"
	"evermos/repository"
)

type TrxItemInput struct {
	ProductID uint `json:"product_id"`
	Kuantitas int  `json:"kuantitas"`
}

type CreateTrxInput struct {
	AlamatKirim uint           `json:"alamat_kirim"`
	MethodBayar string         `json:"method_bayar"`
	DetailTrx   []TrxItemInput `json:"detail_trx"`
}

type TrxDetailResponse struct {
	Trx       models.Trx      `json:"trx"`
	DetailTrx []DetailTrxItem `json:"detail_trx"`
}

type DetailTrxItem struct {
	models.DetailTrx
	LogProduk models.LogProduk `json:"log_produk"`
}

func CreateTransaksi(userID uint, input CreateTrxInput) (*TrxDetailResponse, error) {
	if len(input.DetailTrx) == 0 {
		return nil, errors.New("detail transaksi tidak boleh kosong")
	}

	alamat, err := repository.FindAlamatByID(input.AlamatKirim)
	if err != nil {
		return nil, errors.New("alamat tidak ditemukan")
	}
	if alamat.IdUser != userID {
		return nil, errors.New("alamat bukan milik anda")
	}

	tx := repository.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var totalHarga int
	var detailResponses []DetailTrxItem

	today := time.Now().Format("20060102")
	count := repository.CountTrxToday()
	kodeInvoice := fmt.Sprintf("INV-%s-%03d", today, count+1)

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

	for _, item := range input.DetailTrx {
		produk, err := repository.FindProdukByID(item.ProductID)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("produk dengan ID %d tidak ditemukan", item.ProductID)
		}

		if produk.Stok < item.Kuantitas {
			tx.Rollback()
			return nil, fmt.Errorf("stok produk '%s' tidak cukup (sisa: %d, diminta: %d)",
				produk.NamaProduk, produk.Stok, item.Kuantitas)
		}

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

		if err := repository.UpdateStokProduk(tx, produk.ID, item.Kuantitas); err != nil {
			tx.Rollback()
			return nil, errors.New("gagal mengurangi stok produk")
		}

		hargaItem := produk.HargaKonsumen * item.Kuantitas
		totalHarga += hargaItem

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

	trx.HargaTotal = totalHarga
	if err := tx.Model(&trx).Update("harga_total", totalHarga).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("gagal mengupdate total harga")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, errors.New("gagal menyimpan transaksi")
	}

	return &TrxDetailResponse{
		Trx:       trx,
		DetailTrx: detailResponses,
	}, nil
}

func GetAllTransaksi(userID uint, limit, offset int) ([]models.Trx, int64, error) {
	return repository.GetAllTrxByUserID(userID, limit, offset)
}

func GetDetailTransaksi(userID uint, trxID uint) (*TrxDetailResponse, error) {
	trx, err := repository.FindTrxByID(trxID)
	if err != nil {
		return nil, errors.New("transaksi tidak ditemukan")
	}

	if trx.IdUser != userID {
		return nil, errors.New("transaksi bukan milik anda")
	}

	details, err := repository.GetDetailTrxByTrxID(trxID)
	if err != nil {
		return nil, errors.New("gagal mengambil detail transaksi")
	}

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
