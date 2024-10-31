package services

import (
	"fmt"
	"time"
)

type Log struct {
	UserName string    `csv:"UserName"`
	Message  string    `csv:"Message"`
	Time     time.Time `csv:"Time"`
	Read     bool      `csv:"Read"`
	IsError  bool      `csv:"IsError"`
	LogLevel int       `csv:"LogLevel"`
}

type CachedLog struct {
	UserName string
	Message  string
	Time     time.Time
	Read     bool
	IsError  bool
	LogLevel int
	ID       int
}

func (l *Log) UnmarshalCSV(csv string) error {
	// Adjust the format to match the one in the CSV
	parsedTime, err := time.Parse("Mon Jan 02 2006 15:04:05 GMT-0700 (MST)", csv)
	if err != nil {
		return fmt.Errorf("failed to parse time: %v", err)
	}
	l.Time = parsedTime
	return nil
}
