package handlers

import (
	"SuperSiteLogsApp/services"
	"fmt"
	"sort"
)

const PAGE_SIZE = 100

type Logs struct{}

// GetLogs returns a list of log entries
func (l *Logs) GetLogs(page int) LogsOperationResult {
	logsCache, err := services.GetLogCache()
	if err != nil {
		return LogsOperationResult{
			Success: false,
			Message: err.Error(),
			Logs:    []services.CachedLog{},
		}
	}

	cachedLogArray := logsCache.GetFilteredLogs()
	sort.Slice(cachedLogArray, func(i, j int) bool {
		if cachedLogArray[i].Time.Equal(cachedLogArray[j].Time) {
			// If the time is the same, use ID as a tiebreaker (ascending order)
			return cachedLogArray[i].ID < cachedLogArray[j].ID
		}
		// Otherwise, sort by time (descending order)
		return cachedLogArray[i].Time.After(cachedLogArray[j].Time)
	})

	totalLogs := len(cachedLogArray)
	if totalLogs == 0 {
		return LogsOperationResult{
			Success:    true,
			Message:    "",
			Logs:       []services.CachedLog{},
			TotalCount: 0,
			IsLastPage: true,
		}
	}

	isLastPage := (page+1)*PAGE_SIZE > totalLogs
	var endIndex int
	if !isLastPage {
		endIndex = (page + 1) * PAGE_SIZE
	} else {
		endIndex = totalLogs
	}

	return LogsOperationResult{
		Success:    true,
		Message:    "",
		Logs:       cachedLogArray[page*PAGE_SIZE : endIndex],
		TotalCount: totalLogs,
		IsLastPage: isLastPage,
	}
}

func (l *Logs) SetShowByType(showBy services.LogsToShow) OperationResult {
	appSettings, err := services.GetAppSettings()
	if err != nil {
		return OperationResult{
			Success: false,
			Message: err.Error(),
		}
	}

	appSettings.LogsToShow = showBy
	err = services.SaveAppSettings(appSettings)
	if err != nil {
		return OperationResult{
			Success: false,
			Message: err.Error(),
		}
	}

	return OperationResult{
		Success: true,
	}
}

func (l *Logs) SetLogLevel(logLevel services.LogLevel) OperationResult {
	appSettings, err := services.GetAppSettings()
	if err != nil {
		return OperationResult{
			Success: false,
			Message: err.Error(),
		}
	}

	appSettings.CurrentLogLevel = logLevel
	err = services.SaveAppSettings(appSettings)
	if err != nil {
		return OperationResult{
			Success: false,
			Message: err.Error(),
		}
	}

	return OperationResult{
		Success: true,
	}
}

func (l *Logs) GetLogLevel() int {
	appSettings, err := services.GetAppSettings()
	if err != nil {
		return int(services.AllLogLevels)
	}

	return int(appSettings.CurrentLogLevel)
}

func (l *Logs) GetShowByType() services.LogsToShow {
	appSettings, err := services.GetAppSettings()
	if err != nil {
		return services.ShowAll
	}

	return appSettings.LogsToShow
}

func (l *Logs) MarkAsRead(id int) OperationResult {
	logCache, err := services.GetLogCache()
	if err != nil {
		return OperationResult{
			Success: false,
			Message: err.Error(),
		}
	}

	log := logCache.Logs[id]
	err = services.SetLogAsRead(log)
	if err != nil {
		return OperationResult{
			Success: false,
			Message: err.Error(),
		}
	}

	return OperationResult{
		Success: true,
	}
}

func (l *Logs) MarkAllAsRead() OperationResult {
	logCache, err := services.GetLogCache()
	if err != nil {
		return OperationResult{
			Success: false,
			Message: err.Error(),
		}
	}

	for _, log := range logCache.GetFilteredLogs() {
		err = services.SetLogAsRead(log)
		if err != nil {
			return OperationResult{
				Success: false,
				Message: fmt.Sprintf("Error with log %d: %s", log.ID, err.Error()),
			}
		}
	}

	return OperationResult{
		Success: true,
	}
}
