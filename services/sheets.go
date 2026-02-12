package services

import (
	"context"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func SaveToGoogleSheet(ctx context.Context, data map[string]any) error {
	credentialsJson, err := os.ReadFile("./credentials.json")
	if err != nil {
		return err
	}

	service, err := sheets.NewService(ctx, option.WithAuthCredentialsJSON(option.AuthorizedUser, credentialsJson))

	if err != nil {
		return err
	}

	values := [][]any{}

	value := []any{}
	header := []any{}
	for k, v := range data {
		header = append(header, k)
		value = append(value, v)
	}
	values = append(values, header)
	values = append(values, value)

	sheetId := os.Getenv("SHEET_ID")

	_, err = service.Spreadsheets.Values.Append(sheetId, "Sheet1", &sheets.ValueRange{
		Values: values,
	}).ValueInputOption("RAW").Do()

	if err != nil {
		return err
	}
	return nil
}
