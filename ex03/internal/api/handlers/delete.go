package handlers

import (
	"net/http"

	"ex03/internal/storage"
)

// Delete ...
// Приниамет хранилище
// Удаляет данные по ключу
// Возвращает http.HandlerFunc
func Delete(store storage.I) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
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

		// Валидируем параметры
		if key == "" {
			SendResponse(w, http.StatusBadRequest, Response{Error: "validate error"})
			return
		}

		// Удаляем данные из хранилища
		ok := store.Delete(key)
		if !ok {
			SendResponse(w, http.StatusBadRequest, Response{Error: "key not found"})
			return
		}

		// Возвращаем результат
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		SendResponse(w, http.StatusOK, Response{Result: "event delete successful!"})
	}
}
