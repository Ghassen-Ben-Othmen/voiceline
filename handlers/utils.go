package handlers

import (
	"path"
	"time"
)

func generateAudioName(originalName string) string {
	ext := path.Ext(originalName)
	baseFilename := originalName[:len(originalName)-len(ext)]

	return baseFilename + time.Now().Format(time.DateTime) + ext
}
