package controllers

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
    "intern_backend/output" // Menggunakan package output Anda yang sudah ada
)

// UploadFile adalah controller untuk menangani upload file ke Cloudinary
func UploadFile(c *fiber.Ctx) error {
	// 1. Ambil file dari form request
	fileHeader, err := c.FormFile("file") // "file" adalah nama field di form frontend Anda
	if err != nil {
		return output.GetError(c, http.StatusBadRequest, "File is required")
	}

	// 2. Buka file
	file, err := fileHeader.Open()
	if err != nil {
		return output.GetError(c, http.StatusInternalServerError, "Cannot open file")
	}
	defer file.Close()

	// 3. Panggil fungsi helper untuk upload ke Cloudinary
	uploadResult, err := uploadToCloudinary(file, fileHeader.Filename)
	if err != nil {
		return output.GetError(c, http.StatusInternalServerError, err.Error())
	}

	// Berhasil! Kembalikan URL menggunakan format respons Anda
	return output.GetSuccess(c, "File uploaded successfully", fiber.Map{
		"url": uploadResult.SecureURL, // URL ini yang Anda simpan ke MongoDB
	})
}

// uploadToCloudinary adalah fungsi helper internal untuk berinteraksi dengan API Cloudinary
func uploadToCloudinary(file interface{}, fileName string) (*uploader.UploadResult, error) {
	// Ambil kredensial dari .env
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Lakukan upload dengan nama file asli sebagai Public ID
	uploadResult, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: fileName,
	})
	if err != nil {
		log.Printf("Failed to upload file to Cloudinary: %v\n", err)
		return nil, err
	}

	return uploadResult, nil
}