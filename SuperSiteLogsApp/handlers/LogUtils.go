package handlers

import (
	"SuperSiteLogsApp/services"
)

type LogsUtils struct{}

type CountResult struct {
	Errors       int    `json:"errors"`
	Messages     int    `json:"messages"`
	Success      bool   `json:"success"`
	ErrorMessage string `json:"errorMessage"`
}

func (lu *LogsUtils) GetErrorAndMessageCount() CountResult {
	errors := 0
	messages := 0

	lc, err := services.GetLogCache()
	if err != nil {
		return CountResult{
			Errors:       0,
			Messages:     0,
			Success:      false,
			ErrorMessage: err.Error(),
		}
	}

	logs := lc.GetFilteredLogs()

	for _, log := range logs {
		if log.IsError {
			errors++
			continue
		}

		messages++
	}

	return CountResult{
		Errors:   errors,
		Messages: messages,
	}
}
