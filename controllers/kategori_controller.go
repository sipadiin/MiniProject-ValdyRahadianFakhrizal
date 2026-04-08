package controllers

import (
	"strconv"

	"evermos/services"
	"evermos/utils"

	"github.com/gofiber/fiber/v2"
)

func GetAllKategori(c *fiber.Ctx) error {
	page, limit, offset := utils.GetPaginationParams(c)
	nama := c.Query("nama", "")

	res, total, err := services.ListKategori(limit, offset, nama)
	if err != nil {
		return utils.InternalError(c, "Gagal mengambil kategori", err.Error())
	}
	return utils.Success(c, "Daftar kategori", utils.BuildPagination(res, page, limit, total))
}

func GetKategoriByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	res, err := services.DetailKategori(uint(id))
	if err != nil {
		return utils.NotFound(c, "Kategori tidak ditemukan", nil)
	}
	return utils.Success(c, "Detail kategori", res)
}

func CreateKategori(c *fiber.Ctx) error {
	var input struct {
		NamaKategori string `json:"nama_kategori"`
	}
	c.BodyParser(&input)

	res, err := services.AddKategori(input.NamaKategori)
	if err != nil {
		return utils.InternalError(c, "Gagal simpan", err.Error())
	}
	return utils.Created(c, "Kategori dibuat", res)
}

func UpdateKategori(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var input struct {
		NamaKategori string `json:"nama_kategori"`
	}
	c.BodyParser(&input)

	res, err := services.EditKategori(uint(id), input.NamaKategori)
	if err != nil {
		return utils.NotFound(c, "Kategori tidak ditemukan", err.Error())
	}
	return utils.Success(c, "Kategori diupdate", res)
}

func DeleteKategori(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := services.RemoveKategori(uint(id)); err != nil {
		return utils.NotFound(c, "Kategori tidak ditemukan", err.Error())
	}
	return utils.Success(c, "Kategori dihapus", nil)
}
