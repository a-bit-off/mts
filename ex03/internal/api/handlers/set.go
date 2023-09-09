package handlers

import (
	"net/http"
	"time"

	"ex03/internal/storage"
)

// Set ...
// Приниамет хранилище
// Добавляет в хранилище данные
// Возвращает http.HandlerFunc
func Set(store storage.I) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			SendResponse(w, http.StatusMethodNotAllowed, Response{Error: "method error"})
			return
		}

		// Получаем параметры из тела запроса
		err := r.ParseForm()
		if err != nil {
			SendResponse(w, http.StatusBadRequest, Response{Error: "parse error"})
			return
		}

		key := r.FormValue("key")
		value := r.FormValue("value")

		var ttl time.Time
		ttlStr := r.FormValue("ttl")
		if ttlStr != "" {
			dur, err := time.ParseDuration(ttlStr)
			if err != nil {
				SendResponse(w, http.StatusBadRequest, Response{Error: "parse ttl error"})
				return
			}
			ttl = time.Now().Add(dur)
		} else {
			ttl = storage.Unlimited
		}

		// Валидируем параметры
		if key == "" || value == "" {
			SendResponse(w, http.StatusBadRequest, Response{Error: "validate error"})
			return
		}

		// Добавляем новые данные в хранилище
		store.Set(key, value, ttl)

		// Возвращаем результат
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		SendResponse(w, http.StatusOK, Response{Result: "event set successful!"})
	}
}
