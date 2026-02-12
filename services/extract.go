package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

var ErrDataExtractionFailed = errors.New("sales data extraction failed")

func ExtractSalesData(ctx context.Context, transcript string) (map[string]any, error) {
	prompt := getPromptForDataExtraction(transcript)

	body := map[string]any{
		"model": "gpt-4o-mini",
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
		"response_format": map[string]string{
			"type": "json_object",
		},
	}

	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		OPEN_AI_COMPLETIONS_API_URL,
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return nil, err
	}

	openAIKey := os.Getenv("OPEN_AI_API_KEY")
	req.Header.Set("Authorization", "Bearer "+openAIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, ErrDataExtractionFailed
	}

	defer resp.Body.Close()

	var raw struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}

	var sales map[string]any
	if err := json.Unmarshal([]byte(raw.Choices[0].Message.Content), &sales); err != nil {
		return nil, err
	}

	return sales, nil
}
