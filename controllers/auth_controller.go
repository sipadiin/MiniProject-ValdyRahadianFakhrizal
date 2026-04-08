package controllers

import (
	"evermos/middleware"
	"evermos/services"
	"evermos/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func Register(c *fiber.Ctx) error {
	var input services.RegisterInput
	if err := c.BodyParser(&input); err != nil {
		return utils.BadRequest(c, "Bad Request", "Format body tidak valid")
	}

	if err := validate.Struct(&input); err != nil {
		return utils.BadRequest(c, "Bad Request", err.Error())
	}

	user, err := services.Register(input)
	if err != nil {
		return utils.BadRequest(c, "Register gagal", err.Error())
	}

	token, err := middleware.GenerateToken(user.ID, user.IsAdmin)
	if err != nil {
		return utils.InternalError(c, "Gagal membuat token", err.Error())
	}

	return utils.Created(c, "Register berhasil", fiber.Map{
		"token": token,
		"user":  user,
	})
}

func Login(c *fiber.Ctx) error {
	var input services.LoginInput
	if err := c.BodyParser(&input); err != nil {
		return utils.BadRequest(c, "Bad Request", "Format body tidak valid")
	}

	if err := validate.Struct(&input); err != nil {
		return utils.BadRequest(c, "Bad Request", err.Error())
	}

	user, err := services.Login(input)
	if err != nil {
		return utils.Unauthorized(c, "Login gagal", err.Error())
	}

	token, err := middleware.GenerateToken(user.ID, user.IsAdmin)
	if err != nil {
		return utils.InternalError(c, "Gagal membuat token", err.Error())
	}

	return utils.Success(c, "Login berhasil", fiber.Map{
		"token": token,
		"user":  user,
	})
}
