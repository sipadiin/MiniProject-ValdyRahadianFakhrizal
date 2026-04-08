package controllers

import (
	"strconv"

	"evermos/middleware"
	"evermos/models"
	"evermos/services"
	"evermos/utils"

	"github.com/gofiber/fiber/v2"
)

func GetAlamat(c *fiber.Ctx) error {
	claims := c.Locals("user").(*middleware.Claims)
	page, limit, offset := utils.GetPaginationParams(c)

	alamats, total, err := services.GetAllMyAlamat(claims.UserID, limit, offset)
	if err != nil {
		return utils.InternalError(c, "Gagal mengambil alamat", err.Error())
	}
	return utils.Success(c, "Daftar alamat", utils.BuildPagination(alamats, page, limit, total))
}

func CreateAlamat(c *fiber.Ctx) error {
	claims := c.Locals("user").(*middleware.Claims)
	var alamat models.Alamat
	if err := c.BodyParser(&alamat); err != nil {
		return utils.BadRequest(c, "Format salah", nil)
	}

	res, err := services.CreateNewAlamat(claims.UserID, alamat)
	if err != nil {
		return utils.InternalError(c, "Gagal simpan alamat", err.Error())
	}
	return utils.Created(c, "Alamat berhasil ditambah", res)
}

func UpdateAlamat(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	claims := c.Locals("user").(*middleware.Claims)
	var input models.Alamat
	if err := c.BodyParser(&input); err != nil {
		return utils.BadRequest(c, "Format salah", nil)
	}

	res, err := services.UpdateMyAlamat(uint(id), claims.UserID, input)
	if err != nil {
		return utils.Forbidden(c, "Gagal update", err.Error())
	}
	return utils.Success(c, "Alamat diupdate", res)
}

func DeleteAlamat(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	claims := c.Locals("user").(*middleware.Claims)

	if err := services.DeleteMyAlamat(uint(id), claims.UserID); err != nil {
		return utils.Forbidden(c, "Gagal hapus", err.Error())
	}
	return utils.Success(c, "Alamat dihapus", nil)
}
