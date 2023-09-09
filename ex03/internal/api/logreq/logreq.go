package logreq

import (
	"log"
	"net/http"
)

// LogRequest ...
// middleware для логирования запросов
func LogRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
