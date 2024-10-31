package services

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/gocarina/gocsv"
)

func FetchLogs(date time.Time) ([]Log, error) {
	apiKey, err := GetKey()
	if err != nil {
		return nil, fmt.Errorf("failed to get json key file")
	}

	displayDate := date.Format(time.RFC3339)
	url := fmt.Sprintf("https://script.google.com/macros/s/AKfycbxIPRtyefSe1OcFG1-OqKBWTLJyXCXklzLu6aqqRQB0OeZFSSvOgBaHW4emA9KpNTY/exec?key=%s&date=%s", url.QueryEscape(apiKey), url.QueryEscape(displayDate))

	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch logs: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	if err := saveCSVToConfig(body); err != nil {
		return nil, err
	}

	var logs []Log
	if err := gocsv.Unmarshal(bytes.NewReader(body), &logs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal CSV: %v", err)
	}

	return logs, nil
}

func saveCSVToConfig(body []byte) error {
	// Ensure the config directory exists
	configDir := "config"
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %v", err)
	}

	// Define the path to save the file
	csvFilePath := filepath.Join(configDir, "debug_logs.csv")

	// Write the body to the file
	if err := os.WriteFile(csvFilePath, body, 0644); err != nil {
		return fmt.Errorf("failed to write CSV to file: %v", err)
	}

	fmt.Printf("CSV saved to %s\n", csvFilePath)
	return nil
}
