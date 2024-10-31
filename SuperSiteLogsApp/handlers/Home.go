package handlers

import (
	"time"
)

type Home struct{}

// GetDate returns the current date
func (a *Home) GetDate() string {
	return time.Now().Format("2006-01-02 03:04:05PM")
}
