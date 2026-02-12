package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

var ErrTranscriptionFailed = errors.New("transcription failed")

func TranscribeAudio(ctx context.Context, filePath string) (string, error) {
	file, err := os.ReadFile(filePath)

	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		return "", err
	}

	io.Copy(part, bytes.NewReader(file))

	writer.WriteField("model", OPEN_AI_MODEL)
	writer.Close()

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		OPEN_AI_TRANS_API_URL,
		&b,
	)
	if err != nil {
		return "", err
	}

	openAiApiKey := os.Getenv("OPEN_AI_API_KEY")
	req.Header.Set("Authorization", "Bearer "+openAiApiKey)
	req.Header.Set("Content-type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", ErrTranscriptionFailed
	}

	defer resp.Body.Close()

	var result struct {
		Text string `json:"text"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Text, nil

}
