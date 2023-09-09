package config

import "time"

// Config ...
// Структура с данными конфигурации сервиса
type Config struct {
	Port        string
	RPSLimit    int
	RPSDuration time.Duration
}

// NewConfig ...
// Конструктор для структуры Config
func NewConfig() *Config {
	return &Config{
		Port:        "8080",
		RPSLimit:    10,
		RPSDuration: time.Second * 1,
	}
}
