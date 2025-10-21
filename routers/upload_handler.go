package routers

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
)

// Handler untuk endpoint upload baru Anda
func HandleFileUpload(c *fiber.Ctx) error {
	// 1. Ambil file dari form request
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "file upload failed"})
	}

	// 2. Buka file
	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "cannot open file"})
	}
	defer file.Close()

	// 3. Panggil fungsi untuk upload ke Cloudinary
	uploadResult, err := uploadToCloudinary(file, fileHeader.Filename)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Berhasil! Kembalikan URL aman yang diberikan Cloudinary
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "File uploaded successfully to Cloudinary!",
		"url":     uploadResult.SecureURL, // URL ini yang Anda simpan ke MongoDB
	})
}

// Fungsi inti untuk berinteraksi dengan API Cloudinary
func uploadToCloudinary(file interface{}, fileName string) (*uploader.UploadResult, error) {
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uploadResult, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: fileName,
	})
	if err != nil {
		log.Printf("Failed to upload file to Cloudinary: %v\n", err)
		return nil, err
	}

	return uploadResult, nil
}