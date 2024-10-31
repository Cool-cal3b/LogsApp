package services

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gocarina/gocsv"
)

type LogsToShow int

const (
	ShowAll LogsToShow = iota
	ShowNotRead
)

type LogLevel int

const (
	Message LogLevel = iota
	Warning
	Urgent
	AllLogLevels
)

type AppSettings struct {
	LastSyncDate         time.Time
	LogsToShow           LogsToShow
	CurrentlySyncingLogs bool
	CurrentLogLevel      LogLevel
}

var ConfigDir string
var settingsFilePath string
var LogDatabaseFile string

func init() {
	ConfigDir = "./config"
	settingsFilePath = filepath.Join(ConfigDir, "appSettings.csv")
	LogDatabaseFile = filepath.Join(ConfigDir, "logDatabase.csv")
}

func GetAppSettings() (*AppSettings, error) {
	if err := os.MkdirAll(ConfigDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create config directory: %v", err)
	}

	settingsFile, err := os.ReadFile(settingsFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return &AppSettings{
				LastSyncDate:         time.Date(2024, time.August, 20, 0, 0, 0, 0, time.UTC),
				LogsToShow:           ShowAll,
				CurrentlySyncingLogs: false,
			}, nil
		}

		return nil, err
	}

	var appSettings []*AppSettings
	if err = gocsv.Unmarshal(bytes.NewReader(settingsFile), &appSettings); err != nil {
		return nil, fmt.Errorf("failed to unmarshal CSV: %v", err)
	}

	if len(appSettings) == 0 {
		return nil, fmt.Errorf("settings file is empty")
	}

	return appSettings[0], nil
}

func SaveAppSettings(appSettings *AppSettings) error {
	csvContent, err := gocsv.MarshalString([]*AppSettings{appSettings})
	if err != nil {
		return err
	}

	err = os.WriteFile(settingsFilePath, []byte(csvContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write app settings to file: %v", err)
	}

	return nil
}
