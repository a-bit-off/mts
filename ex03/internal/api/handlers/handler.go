package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response ...
// Структура для возвращаемых значений
type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

// SendResponse ...
// Оборачивает возвращаемое значение в структуру response
func SendResponse(w http.ResponseWriter, statusCode int, responseMessage Response) {
	response, err := json.Marshal(responseMessage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(statusCode)
	w.Write(response)
	log.Println(string(response))
}
