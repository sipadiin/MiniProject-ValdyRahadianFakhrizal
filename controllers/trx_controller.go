package controllers

import (
	"strconv"

	"evermos/middleware"
	"evermos/services"
	"evermos/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateTrx(c *fiber.Ctx) error {
	claims := c.Locals("user").(*middleware.Claims)

	var input services.CreateTrxInput
	if err := c.BodyParser(&input); err != nil {
		return utils.BadRequest(c, "Bad Request", "Format body tidak valid")
	}

	if len(input.DetailTrx) == 0 {
		return utils.BadRequest(c, "Bad Request", "Detail transaksi tidak boleh kosong")
	}

	res, err := services.CreateTransaksi(claims.UserID, input)
	if err != nil {
		return utils.BadRequest(c, "Gagal membuat transaksi", err.Error())
	}

	return utils.Created(c, "Transaksi berhasil dibuat", res)
}

func GetAllTrx(c *fiber.Ctx) error {
	claims := c.Locals("user").(*middleware.Claims)

	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	page, _ := strconv.Atoi(c.Query("page", "1"))
	offset := (page - 1) * limit

	trxs, total, err := services.GetAllTransaksi(claims.UserID, limit, offset)
	if err != nil {
		return utils.InternalError(c, "Gagal mengambil data transaksi", err.Error())
	}

	return utils.Success(c, "Daftar Transaksi", fiber.Map{
		"data":  trxs,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func GetTrxByID(c *fiber.Ctx) error {
	claims := c.Locals("user").(*middleware.Claims)

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.BadRequest(c, "ID tidak valid", nil)
	}

	res, err := services.GetDetailTransaksi(claims.UserID, uint(id))
	if err != nil {
		return utils.NotFound(c, err.Error(), nil)
	}

	return utils.Success(c, "Detail Transaksi", res)
}
