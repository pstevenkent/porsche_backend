package controllers

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings" // <-- Pastikan import ini ada
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
	"intern_backend/output"
)

// UploadFile adalah controller untuk menangani upload file ke Cloudinary
func UploadFile(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return output.GetError(c, http.StatusBadRequest, "File is required")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return output.GetError(c, http.StatusInternalServerError, "Cannot open file")
	}
	defer file.Close()

	uploadResult, err := uploadToCloudinary(file, fileHeader.Filename)
	if err != nil {
		return output.GetError(c, http.StatusInternalServerError, err.Error())
	}

	return output.GetSuccess(c, "File uploaded successfully", fiber.Map{
		"url": uploadResult.SecureURL,
	})
}

// uploadToCloudinary adalah fungsi helper internal untuk berinteraksi dengan API Cloudinary
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

	// --- INI ADALAH LOGIKA BARU YANG MEMPERBAIKI PDF ---
	var resourceType string
	if strings.HasSuffix(strings.ToLower(fileName), ".pdf") {
		resourceType = "raw" // Perlakukan PDF sebagai file mentah
	} else {
		resourceType = "image" // Perlakukan file lain sebagai gambar
	}

	uploadParams := uploader.UploadParams{
		PublicID:     fileName,
		ResourceType: resourceType, // Set resource type di sini
	}
	// --- AKHIR DARI LOGIKA BARU ---

	// Lakukan upload dengan parameter yang sudah disesuaikan
	uploadResult, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		log.Printf("Failed to upload file to Cloudinary: %v\n", err)
		return nil, err
	}

	return uploadResult, nil
}