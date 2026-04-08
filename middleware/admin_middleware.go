package middleware

import (
	"evermos/utils"

	"github.com/gofiber/fiber/v2"
)

func AdminMiddleware(c *fiber.Ctx) error {
	claims, ok := c.Locals("user").(*Claims)
	if !ok || !claims.IsAdmin {
		return utils.Forbidden(c, "Akses Terlarang", "Hanya admin yang diperbolehkan melakukan aksi ini")
	}
	return c.Next()
}
