package log

import (
	"fmt"
	"time"

	"ex01/pkg/mylogger/collor"
	"ex01/pkg/mylogger/level"
)

// Log ...
// Структура лога
type Log struct {
	id           int
	level        level.Level
	creationTime time.Time
	message      string
}

// NewLog ...
// Конструктор для структуры Log
func NewLog(id int, level level.Level, message string) Log {
	return Log{
		id:           id,
		level:        level,
		creationTime: time.Now(),
		message:      message,
	}
}

// LogPretty ...
// Преобразование структуры лога в красивую строку
func (l *Log) LogPretty() string {
	logLevel := ""
	switch l.level {
	case level.Debug:
		logLevel = collor.ColorizeString("DEBUG", "blue")
	case level.Info:
		logLevel = collor.ColorizeString("INFO", "green")
	case level.Warn:
		logLevel = collor.ColorizeString("WARN", "yellow")
	case level.Error:
		logLevel = collor.ColorizeString("ERROR", "red")
	}
	return fmt.Sprintf("{\"time\":\"%s\","+
		"\"id\":\"%d\","+
		"\"level\":\"%s\","+
		"\"msg\":\"%s\"}", l.creationTime.Format("2006-01-02T15:04:05.999999-07:00"), l.id, logLevel, l.message)
}
