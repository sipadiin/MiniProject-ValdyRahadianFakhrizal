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

	// 1. Validasi Ekstensi
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return "", fmt.Errorf("format file tidak didukung (hanya jpg, jpeg, png)")
	}

	// 2. Buat folder uploads jika belum ada
	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}

	// 3. Generate Nama Unik
	uniqueName := fmt.Sprintf("%d-%s", time.Now().UnixNano(), strings.ReplaceAll(file.Filename, " ", "_"))
	filePath := filepath.Join(uploadDir, uniqueName)

	// 4. Simpan File
	if err := c.SaveFile(file, filePath); err != nil {
		return "", err
	}

	return uniqueName, nil
}
