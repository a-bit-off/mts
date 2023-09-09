package mylogger

import (
	"fmt"
	"io"

	"ex01/pkg/mylogger/level"
	"ex01/pkg/mylogger/log"
)

// Logger ...
// Основная структура для хранения логов
type Logger struct {
	output io.Writer
	totalD int
	log    log.Log
}

// NewLogger ...
// Конструктор для структуры Logger
func NewLogger(writer io.Writer) *Logger {
	return &Logger{
		output: writer,
	}
}

// Debug ...
// Лог для дебага
// Принимает сообщение
func (l *Logger) Debug(msg string) {
	l.logging(msg, level.Debug)
}

// Info ...
// Лог для информации
// Принимает сообщение
func (l *Logger) Info(msg string) {
	l.logging(msg, level.Info)
}

// Warn ...
// Лог для предупреждения
// Принимает сообщение
func (l *Logger) Warn(msg string) {
	l.logging(msg, level.Warn)
}

// Error ...
// Лог для ошибки
// Принимает сообщение
func (l *Logger) Error(msg string) {
	l.logging(msg, level.Error)
}

// logging ...
// Вывод лога в io.Writer
// Принимает сообщение и уровень
func (l *Logger) logging(msg string, lvl level.Level) {
	l.log = log.NewLog(l.getID(), lvl, msg)
	fmt.Fprintln(l.output, l.log.LogPretty())
}

// getID ...
// Получаем уникальный id
func (l *Logger) getID() int {
	id := l.totalD
	l.totalD++
	return id
}
