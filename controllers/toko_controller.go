package controllers

import (
	"strconv"

	"evermos/middleware"
	"evermos/repository"
	"evermos/services"
	"evermos/utils"

	"github.com/gofiber/fiber/v2"
)

func GetAllToko(c *fiber.Ctx) error {
	page, limit, offset := utils.GetPaginationParams(c)
	nama := c.Query("nama", "")

	tokos, total, err := services.GetAllToko(page, limit, nama)
	if err != nil {
		return utils.InternalError(c, "Gagal mengambil data toko", err.Error())
	}

	_ = offset
	return utils.Success(c, "Berhasil mengambil semua toko", utils.BuildPagination(tokos, page, limit, total))
}

func GetTokoByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id_toko"))
	toko, err := services.GetTokoByID(uint(id))
	if err != nil {
		return utils.NotFound(c, "Toko tidak ditemukan", err.Error())
	}
	return utils.Success(c, "Berhasil mengambil detail toko", toko)
}

func GetMyToko(c *fiber.Ctx) error {
	claims := c.Locals("user").(*middleware.Claims)
	toko, err := repository.FindTokoByUserID(claims.UserID)
	if err != nil {
		return utils.NotFound(c, "Toko tidak ditemukan", err.Error())
	}
	return utils.Success(c, "Succeed", toko)
}

func UpdateToko(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id_toko"))
	userLocal := c.Locals("user")
	claims, ok := userLocal.(*middleware.Claims)

	if !ok || claims == nil {
		return utils.Unauthorized(c, "Unauthorized", "Gagal mendapatkan identitas user")
	}

	var input services.UpdateTokoInput
	if err := c.BodyParser(&input); err != nil {
		return utils.BadRequest(c, "Format input salah", nil)
	}

	toko, err := services.UpdateMyToko(uint(id), claims.UserID, input)
	if err != nil {
		return utils.Forbidden(c, "Gagal update toko", err.Error())
	}

	return utils.Success(c, "Toko berhasil diupdate", toko)
}
