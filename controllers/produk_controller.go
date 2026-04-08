package controllers

import (
	"strconv"

	"evermos/middleware"
	"evermos/models"
	"evermos/repository"
	"evermos/services"
	"evermos/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateProduk(c *fiber.Ctx) error {
	claims := c.Locals("user").(*middleware.Claims)

	var prod models.Produk
	prod.NamaProduk = c.FormValue("nama_produk")
	catID, _ := strconv.Atoi(c.FormValue("id_category"))
	prod.IdCategory = uint(catID)
	hRes, _ := strconv.Atoi(c.FormValue("harga_reseller"))
	prod.HargaReseller = hRes
	hKon, _ := strconv.Atoi(c.FormValue("harga_konsumen"))
	prod.HargaKonsumen = hKon
	stok, _ := strconv.Atoi(c.FormValue("stok"))
	prod.Stok = stok
	prod.Deskripsi = c.FormValue("deskripsi")

	fileName, err := utils.HandleSingleUpload(c, "foto")
	var fotos []string
	if err == nil {
		fotos = append(fotos, fileName)
	}

	res, err := services.CreateNewProduk(claims.UserID, prod, fotos)
	if err != nil {
		return utils.InternalError(c, "Gagal simpan produk", err.Error())
	}

	return utils.Created(c, "Produk berhasil dibuat", res)
}

func GetAllProduk(c *fiber.Ctx) error {
	page, limit, offset := utils.GetPaginationParams(c)

	filters := map[string]interface{}{
		"nama":        c.Query("nama_produk"),
		"id_category": c.Query("category_id"),
		"id_toko":     c.Query("toko_id"),
		"min_harga":   c.Query("min_harga"),
		"max_harga":   c.Query("max_harga"),
	}

	res, total, _ := repository.GetAllProduk(limit, offset, filters)
	return utils.Success(c, "Daftar Produk", utils.BuildPagination(res, page, limit, total))
}

func GetProdukByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.BadRequest(c, "ID tidak valid", nil)
	}

	produk, err := repository.FindProdukByID(uint(id))
	if err != nil {
		return utils.NotFound(c, "Produk tidak ditemukan", err.Error())
	}

	return utils.Success(c, "Detail produk", produk)
}

func UpdateProduk(c *fiber.Ctx) error {
	claims := c.Locals("user").(*middleware.Claims)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.BadRequest(c, "ID tidak valid", nil)
	}

	produk, err := repository.FindProdukByID(uint(id))
	if err != nil {
		return utils.NotFound(c, "Produk tidak ditemukan", err.Error())
	}

	toko, err := repository.FindTokoByID(produk.IdToko)
	if err != nil || toko.UserID != claims.UserID {
		return utils.Forbidden(c, "Forbidden", "Anda tidak memiliki akses ke produk ini")
	}

	if nama := c.FormValue("nama_produk"); nama != "" {
		produk.NamaProduk = nama
	}
	if catID, err := strconv.Atoi(c.FormValue("id_category")); err == nil && catID > 0 {
		produk.IdCategory = uint(catID)
	}
	if hRes, err := strconv.Atoi(c.FormValue("harga_reseller")); err == nil && hRes > 0 {
		produk.HargaReseller = hRes
	}
	if hKon, err := strconv.Atoi(c.FormValue("harga_konsumen")); err == nil && hKon > 0 {
		produk.HargaKonsumen = hKon
	}
	if stok, err := strconv.Atoi(c.FormValue("stok")); err == nil {
		produk.Stok = stok
	}
	if desc := c.FormValue("deskripsi"); desc != "" {
		produk.Deskripsi = desc
	}

	if err := repository.UpdateProduk(produk); err != nil {
		return utils.InternalError(c, "Gagal update produk", err.Error())
	}

	return utils.Success(c, "Produk berhasil diupdate", produk)
}

func DeleteProduk(c *fiber.Ctx) error {
	claims := c.Locals("user").(*middleware.Claims)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.BadRequest(c, "ID tidak valid", nil)
	}

	produk, err := repository.FindProdukByID(uint(id))
	if err != nil {
		return utils.NotFound(c, "Produk tidak ditemukan", err.Error())
	}

	toko, err := repository.FindTokoByID(produk.IdToko)
	if err != nil || toko.UserID != claims.UserID {
		return utils.Forbidden(c, "Forbidden", "Anda tidak memiliki akses ke produk ini")
	}

	if err := repository.DeleteProduk(uint(id)); err != nil {
		return utils.InternalError(c, "Gagal hapus produk", err.Error())
	}

	return utils.Success(c, "Produk berhasil dihapus", nil)
}
