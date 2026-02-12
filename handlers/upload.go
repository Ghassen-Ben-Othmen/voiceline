package handlers

import (
	"log"
	"net/http"
	"path"

	"github.com/G-b-o/voice-line/services"
	"github.com/gin-gonic/gin"
)

var (
	UPLOAD_DIR_PATH = path.Join(".", "uploads")
)

func UploadAudio(ctx *gin.Context) {
	file, err := ctx.FormFile("audio")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save locally
	filename := generateAudioName(file.Filename)
	filePath := path.Join(UPLOAD_DIR_PATH, filename)
	log.Println(filePath)
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Transcribe the audio file
	transcript, err := services.TranscribeAudio(ctx, filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Extract sales data from the transcript
	salesData, err := services.ExtractSalesData(ctx, transcript)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Save the extracted data to Google Sheets
	if err := services.SaveToGoogleSheet(ctx, salesData); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Audio uploaded and processed successfully", "data": salesData})
}
