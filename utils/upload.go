package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func HandleSingleUpload(c *fiber.Ctx, fieldName string) (string, error) {
	file, err := c.FormFile(fieldName)
	if err != nil {
		return "", err
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return "", fmt.Errorf("format file tidak didukung (hanya jpg, jpeg, png)")
	}

	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}

	uniqueName := fmt.Sprintf("%d-%s", time.Now().UnixNano(), strings.ReplaceAll(file.Filename, " ", "_"))
	filePath := filepath.Join(uploadDir, uniqueName)

	if err := c.SaveFile(file, filePath); err != nil {
		return "", err
	}

	return uniqueName, nil
}
