package handlers

import "SuperSiteLogsApp/services"

type LogView struct{}

func (l *LogView) GetLog(logID int) LogViewOperationResult {
	logCache, err := services.GetLogCache()
	if err != nil {
		return LogViewOperationResult{
			Success: false,
			Message: err.Error(),
		}
	}

	log := logCache.Logs[logID]
	go services.SetLogAsRead(log)

	return LogViewOperationResult{
		Success: true,
		Log:     log,
	}
}
