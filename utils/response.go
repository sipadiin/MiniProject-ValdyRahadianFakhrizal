package utils

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

func Success(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Status:  true,
		Message: message,
		Errors:  nil,
		Data:    data,
	})
}

func Created(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(Response{
		Status:  true,
		Message: message,
		Errors:  nil,
		Data:    data,
	})
}

func BadRequest(c *fiber.Ctx, message string, errors interface{}) error {
	return c.Status(fiber.StatusBadRequest).JSON(Response{
		Status:  false,
		Message: message,
		Errors:  errors,
		Data:    nil,
	})
}

func Unauthorized(c *fiber.Ctx, message string, errors interface{}) error {
	return c.Status(fiber.StatusUnauthorized).JSON(Response{
		Status:  false,
		Message: message,
		Errors:  errors,
		Data:    nil,
	})
}

func Forbidden(c *fiber.Ctx, message string, errors interface{}) error {
	return c.Status(fiber.StatusForbidden).JSON(Response{
		Status:  false,
		Message: message,
		Errors:  errors,
		Data:    nil,
	})
}

func NotFound(c *fiber.Ctx, message string, errors interface{}) error {
	return c.Status(fiber.StatusNotFound).JSON(Response{
		Status:  false,
		Message: message,
		Errors:  errors,
		Data:    nil,
	})
}

func InternalError(c *fiber.Ctx, message string, errors interface{}) error {
	return c.Status(fiber.StatusInternalServerError).JSON(Response{
		Status:  false,
		Message: message,
		Errors:  errors,
		Data:    nil,
	})
}
