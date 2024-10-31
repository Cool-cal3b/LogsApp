package services

import (
	"bytes"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

func AddLogsToDB(logs []Log) error {
	if err := os.MkdirAll(ConfigDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %v", err)
	}

	fileExists := true
	if _, err := os.Stat(LogDatabaseFile); os.IsNotExist(err) {
		fileExists = false
	}

	logFile, err := os.OpenFile(LogDatabaseFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to create config directory: %v", err)
	}
	defer logFile.Close()

	if !fileExists {
		if err := gocsv.MarshalFile(logs, logFile); err != nil {
			return fmt.Errorf("failed to marshal logs with headers to CSV: %v", err)
		}
	} else {
		if err := gocsv.MarshalWithoutHeaders(logs, logFile); err != nil {
			return fmt.Errorf("failed to marshal new logs to CSV: %v", err)
		}
	}

	return nil
}

func GetLogsFromDB() ([]*Log, error) {
	if err := os.MkdirAll(ConfigDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create config directory: %v", err)
	}

	logsFile, err := os.ReadFile(LogDatabaseFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []*Log{}, nil
		}

		return nil, err
	}

	var logs []*Log
	if err = gocsv.Unmarshal(bytes.NewReader(logsFile), &logs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal CSV: %v", err)
	}

	if len(logs) == 0 {
		return nil, fmt.Errorf("settings file is empty")
	}

	return logs, nil
}

func SetLogAsRead(log CachedLog) error {
	logCache, err := GetLogCache()
	if err != nil {
		return err
	}

	log.Read = true
	logCache.Logs[log.ID] = log

	file, err := os.ReadFile(LogDatabaseFile)
	if err != nil {
		return err
	}

	var logs []Log
	if err := gocsv.UnmarshalBytes(file, &logs); err != nil {
		return err
	}

	logs[log.ID].Read = true

	csvData, err := gocsv.MarshalBytes(&logs)
	if err != nil {
		return err
	}

	if err := os.WriteFile(LogDatabaseFile, csvData, 0644); err != nil {
		return err
	}

	return nil
}
