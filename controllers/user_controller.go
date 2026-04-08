package controllers

import (
	"evermos/middleware"
	"evermos/services"
	"evermos/utils"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return utils.Unauthorized(c, "Unauthorized", "User ID tidak valid")
	}

	user, err := services.GetMyProfile(userID)
	if err != nil {
		return utils.NotFound(c, "Gagal mengambil profil", err.Error())
	}

	return utils.Success(c, "Berhasil mengambil profil", user)
}

func UpdateUser(c *fiber.Ctx) error {
	claims := c.Locals("user").(*middleware.Claims)
	userID := claims.UserID

	var input services.UpdateUserInput
	if err := c.BodyParser(&input); err != nil {
		return utils.BadRequest(c, "Format input salah", nil)
	}

	user, err := services.UpdateMyProfile(userID, input)
	if err != nil {
		return utils.BadRequest(c, "Gagal update profil", err.Error())
	}

	return utils.Success(c, "Profil berhasil diupdate", user)
}
