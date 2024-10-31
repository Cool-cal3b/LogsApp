package services

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const keyDirName = "config/jsonkey"

type APIKey struct {
	Key string `json:"key"`
}

// SetKeyPath clears the jsonkey folder and stores the provided file
func SetKey(base64Content string, fileName string) error {
	if fileName == "" {
		return errors.New("file name cannot be empty")
	}

	fileContent, err := base64.StdEncoding.DecodeString(base64Content)
	if err != nil {
		return err
	}

	// Ensure the directory exists
	if err := os.MkdirAll(keyDirName, 0755); err != nil {
		return err
	}

	// Clear the directory
	files, err := os.ReadDir(keyDirName)
	if err != nil {
		return err
	}

	for _, file := range files {
		err = os.Remove(filepath.Join(keyDirName, file.Name()))
		if err != nil {
			return err
		}
	}

	// Write the new file
	return os.WriteFile(filepath.Join(keyDirName, fileName), fileContent, 0644)
}

// GetKey retrieves the content of the first file in the jsonkey folder and parses the json file and gets the key
func GetKey() (string, error) {
	// Ensure the directory exists
	if err := os.MkdirAll(keyDirName, 0755); err != nil {
		return "", err
	}

	files, err := os.ReadDir(keyDirName)
	if err != nil {
		return "", err
	}

	if len(files) <= 0 {
		return "", nil
	}

	keyJsonFile, err := os.ReadFile(filepath.Join(keyDirName, files[0].Name()))
	if err != nil {
		return "", fmt.Errorf("error reading key json file: %v", err)
	}

	var apiKey APIKey
	err = json.Unmarshal(keyJsonFile, &apiKey)
	if err != nil {
		return "", fmt.Errorf("error parsing json key file: %v", err)
	}

	return apiKey.Key, nil
}

// GetKeyName retrieves the filename of the first file in the jsonkey folder
func GetKeyName() (string, error) {
	if err := os.MkdirAll(keyDirName, 0755); err != nil {
		return "", err
	}

	files, err := os.ReadDir(keyDirName)
	if err != nil {
		return "", err
	}

	if len(files) > 0 {
		return files[0].Name(), nil
	}

	return "", nil
}
