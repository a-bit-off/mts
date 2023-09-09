package inmemory

import (
	"log"
	"sync"
	"time"
)

type keyValue struct {
	Value     interface{}
	ExpiresAt time.Time
}

// MemoryStorage ...
type MemoryStorage struct {
	mu sync.Mutex

	lastHour int
	data     map[string]keyValue
}

// NewMemoryStorage ...
// Конструктор для структуры MemoryStorage
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		mu:   sync.Mutex{},
		data: make(map[string]keyValue),
	}
}

// Set ...
// Добавляет данные в хранилище
// Принимает ключ, значение и срок хранения данных
func (s *MemoryStorage) Set(key string, value interface{}, ttl time.Time) {
	s.mu.Lock()
	s.data[key] = keyValue{
		Value:     value,
		ExpiresAt: ttl,
	}
	s.mu.Unlock()

	s.clearStorage()
}

// Delete ...
// Удаляем данные из хранилища по ключу
// Принимает ключ
// Возвращает успешность выполнения опреации
func (s *MemoryStorage) Delete(key string) bool {
	if _, ok := s.data[key]; ok {
		s.mu.Lock()
		delete(s.data, key)
		s.mu.Unlock()

		return true
	}
	s.clearStorage()

	return false
}

// clearStorage ...
// очищаем хранилище каждый час (без нагрузки на CPU)
// каждый раз когда, когда происходит запрос к хранилищу, проверяем час запроса,
// если он отличен от s.lastHour - очищаем
func (s *MemoryStorage) clearStorage() {
	currentHour := time.Now().Hour()
	if currentHour > s.lastHour || currentHour == 0 {
		s.lastHour = currentHour

		s.mu.Lock()
		for k, v := range s.data {
			if !time.Now().Before(v.ExpiresAt) {
				log.Println("DELETED key:", k)
				delete(s.data, k)
			}
		}
		s.mu.Unlock()
	}
}
