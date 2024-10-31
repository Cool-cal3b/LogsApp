package handlers

import (
	"SuperSiteLogsApp/services"
	"time"
)

type SyncLogs struct{}

func (s *SyncLogs) SyncLogs() OperationResult {
	appSettings, err := services.GetAppSettings()
	if err != nil || appSettings.CurrentlySyncingLogs {
		return OperationResult{
			Success: false,
			Message: err.Error(),
		}
	}

	appSettings.CurrentlySyncingLogs = true
	services.SaveAppSettings(appSettings)

	newLogs, err := services.FetchLogs(appSettings.LastSyncDate)
	if err == nil {
		err = services.AddLogsToDB(newLogs)
	}

	services.ResetLogsCache()

	if err != nil {
		return OperationResult{
			Success: false,
			Message: err.Error(),
		}
	}

	appSettings.LastSyncDate = time.Now()
	appSettings.CurrentlySyncingLogs = false
	services.SaveAppSettings(appSettings)

	return OperationResult{
		Success: true,
		Message: "",
	}
}

func (s SyncLogs) IsCurrentlyLoading() bool {
	appSettings, err := services.GetAppSettings()
	return err == nil && appSettings.CurrentlySyncingLogs
}
