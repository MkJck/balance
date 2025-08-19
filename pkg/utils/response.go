package utils

import (
	"encoding/json"
	"net/http"
)

// Response представляет стандартную структуру ответа API
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SendJSON отправляет JSON ответ
func SendJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "JSON encoding error", http.StatusInternalServerError)
	}
}

// SendSuccess отправляет успешный ответ
func SendSuccess(w http.ResponseWriter, data interface{}) {
	response := Response{
		Status: "success",
		Data:   data,
	}
	SendJSON(w, http.StatusOK, response)
}

// SendCreated отправляет ответ о создании ресурса
func SendCreated(w http.ResponseWriter, data interface{}) {
	response := Response{
		Status:  "success",
		Message: "Resource created successfully",
		Data:    data,
	}
	SendJSON(w, http.StatusCreated, response)
}

// SendError отправляет ответ с ошибкой
func SendError(w http.ResponseWriter, statusCode int, message string) {
	response := Response{
		Status: "error",
		Error:  message,
	}
	SendJSON(w, statusCode, response)
}

// SendBadRequest отправляет ответ с ошибкой валидации
func SendBadRequest(w http.ResponseWriter, message string) {
	SendError(w, http.StatusBadRequest, message)
}

// SendNotFound отправляет ответ о ненайденном ресурсе
func SendNotFound(w http.ResponseWriter, message string) {
	SendError(w, http.StatusNotFound, message)
}

// SendInternalError отправляет ответ о внутренней ошибке сервера
func SendInternalError(w http.ResponseWriter, message string) {
	SendError(w, http.StatusInternalServerError, message)
}
