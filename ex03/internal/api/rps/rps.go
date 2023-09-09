package rps

import (
	"net/http"
	"sync"
	"time"

	"ex03/internal/api/handlers"
)

// Limiter ...
// для устаноления лимита кол-ва запросов
type Limiter struct {
	mutex       sync.RWMutex
	rpsLimiter  map[string]*time.Ticker
	rpsLimit    int
	rpsDuration time.Duration
}

// NewRPSLimiter ...
// Конструктор для структуры Limiter
func NewRPSLimiter(rpsLimit int, rpsDuration time.Duration) *Limiter {
	return &Limiter{
		mutex:       sync.RWMutex{},
		rpsLimiter:  make(map[string]*time.Ticker),
		rpsLimit:    rpsLimit,
		rpsDuration: rpsDuration / time.Duration(rpsLimit),
	}
}

// LimitIP ...
// Устанавливаем лимит на запросы по ip пользователя
func (l *Limiter) LimitIP(ip string) bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if _, ok := l.rpsLimiter[ip]; !ok {
		l.rpsLimiter[ip] = time.NewTicker(l.rpsDuration)
		go func() {
			<-l.rpsLimiter[ip].C
			l.mutex.Lock()
			delete(l.rpsLimiter, ip)
			l.mutex.Unlock()
		}()
		return true
	}
	return false

}

// RPSLimit ...
// middleware для логирования запросов
func (l *Limiter) RPSLimit(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := "1" // TODO: CHANGE
		if l.LimitIP(ip) {
			next.ServeHTTP(w, r)
		} else {
			handlers.SendResponse(w, http.StatusTooManyRequests, handlers.Response{Result: "Too many requests!"})
		}
	}
}
