package handlers

import "SuperSiteLogsApp/services"

type OperationResult struct {
	Success bool
	Message string
}

type LogsOperationResult struct {
	Success    bool
	Message    string
	Logs       []services.CachedLog
	TotalCount int
	IsLastPage bool
}

type LogViewOperationResult struct {
	Success bool
	Message string
	Log     services.CachedLog
}
