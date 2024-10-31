package services

import (
	"sync"
)

type LogCache struct {
	Logs map[int]CachedLog
	mu   sync.Mutex
}

var instance *LogCache
var once sync.Once

func GetLogCache() (*LogCache, error) {
	var initError error

	once.Do(func() {
		initError = ResetLogsCache()
	})

	return instance, initError
}

func ResetLogsCache() error {
	var err error
	instance, err = getLogCacheFirstTime()
	return err
}

func getLogCacheFirstTime() (*LogCache, error) {
	allLogs, err := GetLogsFromDB()
	if err != nil {
		return nil, err
	}

	lc := &LogCache{
		Logs: make(map[int]CachedLog),
	}

	for i, log := range allLogs {
		lc.AddLog(CachedLog{
			Time:     log.Time,
			Message:  log.Message,
			ID:       i,
			IsError:  log.IsError,
			Read:     log.Read,
			UserName: log.UserName,
			LogLevel: log.LogLevel,
		})
	}

	return lc, nil
}

func (lc *LogCache) AddLog(log CachedLog) {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	lc.Logs[log.ID] = log
}

func (lc *LogCache) MarkAsRead(id int) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	if log, exists := lc.Logs[id]; exists {
		log.Read = true
		lc.Logs[id] = log
	}
}

func (lc *LogCache) GetFilteredLogs() []CachedLog {
	appSettings, err := GetAppSettings()
	if err != nil {
		return []CachedLog{}
	}

	var cachedLogArray []CachedLog
	for _, log := range lc.Logs {
		if appSettings.LogsToShow == ShowNotRead && log.Read {
			continue
		}

		if appSettings.CurrentLogLevel == Message && log.LogLevel != int(Message) {
			continue
		}

		if appSettings.CurrentLogLevel == Warning && log.LogLevel != int(Warning) {
			continue
		}

		if appSettings.CurrentLogLevel == Urgent && log.LogLevel != int(Urgent) {
			continue
		}

		cachedLogArray = append(cachedLogArray, log)
	}

	return cachedLogArray
}
