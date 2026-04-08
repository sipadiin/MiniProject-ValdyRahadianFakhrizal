package utils

import "github.com/gofiber/fiber/v2"

// Response adalah format standar untuk semua API response
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// Success mengirim response sukses
func Success(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  true,
		Message: message,
		Errors:  nil,
		Data:    data,
	})
}

// Created mengirim response 201 (untuk create berhasil)
func Created(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(Response{
		Status:  true,
		Message: message,
		Errors:  nil,
		Data:    data,
	})
}

// BadRequest mengirim response 400
func BadRequest(c *fiber.Ctx, message string, errors interface{}) error {
	return c.Status(fiber.StatusBadRequest).JSON(Response{
		Status:  false,
		Message: message,
		Errors:  errors,
		Data:    nil,
	})
}

// Unauthorized mengirim response 401
func Unauthorized(c *fiber.Ctx, message string, errors interface{}) error {
	return c.Status(fiber.StatusUnauthorized).JSON(Response{
		Status:  false,
		Message: message,
		Errors:  errors,
		Data:    nil,
	})
}

// Forbidden mengirim response 403
func Forbidden(c *fiber.Ctx, message string, errors interface{}) error {
	return c.Status(fiber.StatusForbidden).JSON(Response{
		Status:  false,
		Message: message,
		Errors:  errors,
		Data:    nil,
	})
}

// NotFound mengirim response 404
func NotFound(c *fiber.Ctx, message string, errors interface{}) error {
	return c.Status(fiber.StatusNotFound).JSON(Response{
		Status:  false,
		Message: message,
		Errors:  errors,
		Data:    nil,
	})
}

// InternalError mengirim response 500
func InternalError(c *fiber.Ctx, message string, errors interface{}) error {
	return c.Status(fiber.StatusInternalServerError).JSON(Response{
		Status:  false,
		Message: message,
		Errors:  errors,
		Data:    nil,
	})
}
