package storage

import (
	"time"
)

// Unlimited ...
// Для бессрочного хранения данных
var Unlimited = time.Date(9999, time.December, 31, 23, 59, 59, 0, time.UTC)

// I ...
// Интерфейс I для добавления возможных сценариев хранения (PostgresSQL, SQL и тд.)
type I interface {
	Set(key string, value interface{}, ttl time.Time)
	Delete(key string) bool
}
