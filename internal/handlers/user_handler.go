package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"balance/internal/models"
	"balance/internal/service"
	"balance/pkg/utils"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// CreateUser обрабатывает POST запрос для создания пользователя
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.SendError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Ограничиваем размер тела запроса
	r.Body = http.MaxBytesReader(w, r.Body, 1048576) // 1MB

	// Декодируем JSON
	var req models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendBadRequest(w, "Invalid JSON format")
		return
	}

	// Создаем пользователя
	user, err := h.userService.CreateUser(r.Context(), &req)
	if err != nil {
		if err.Error() == "user with this email already exists" {
			utils.SendError(w, http.StatusConflict, err.Error())
			return
		}
		utils.SendBadRequest(w, err.Error())
		return
	}

	utils.SendCreated(w, user)
}

// GetUser обрабатывает GET запрос для получения пользователя по ID
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.SendError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Извлекаем ID из URL
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		utils.SendBadRequest(w, "User ID is required")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendBadRequest(w, "Invalid user ID")
		return
	}

	// Получаем пользователя
	user, err := h.userService.GetUser(r.Context(), id)
	if err != nil {
		if err.Error() == "user not found" {
			utils.SendNotFound(w, "User not found")
			return
		}
		utils.SendInternalError(w, "Failed to get user")
		return
	}

	utils.SendSuccess(w, user)
}

// GetAllUsers обрабатывает GET запрос для получения всех пользователей
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.SendError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Получаем всех пользователей
	users, err := h.userService.GetAllUsers(r.Context())
	if err != nil {
		utils.SendInternalError(w, "Failed to get users")
		return
	}

	utils.SendSuccess(w, users)
}

// UpdateUser обрабатывает PUT запрос для обновления пользователя
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		utils.SendError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Извлекаем ID из URL
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		utils.SendBadRequest(w, "User ID is required")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendBadRequest(w, "Invalid user ID")
		return
	}

	// Ограничиваем размер тела запроса
	r.Body = http.MaxBytesReader(w, r.Body, 1048576) // 1MB

	// Декодируем JSON
	var req models.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendBadRequest(w, "Invalid JSON format")
		return
	}

	// Обновляем пользователя
	user, err := h.userService.UpdateUser(r.Context(), id, &req)
	if err != nil {
		if err.Error() == "user not found" {
			utils.SendNotFound(w, "User not found")
			return
		}
		if err.Error() == "user with this email already exists" {
			utils.SendError(w, http.StatusConflict, err.Error())
			return
		}
		utils.SendBadRequest(w, err.Error())
		return
	}

	utils.SendSuccess(w, user)
}

// DeleteUser обрабатывает DELETE запрос для удаления пользователя
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		utils.SendError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Извлекаем ID из URL
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		utils.SendBadRequest(w, "User ID is required")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendBadRequest(w, "Invalid user ID")
		return
	}

	// Удаляем пользователя
	if err := h.userService.DeleteUser(r.Context(), id); err != nil {
		if err.Error() == "user not found" {
			utils.SendNotFound(w, "User not found")
			return
		}
		utils.SendInternalError(w, "Failed to delete user")
		return
	}

	utils.SendSuccess(w, map[string]string{"message": "User deleted successfully"})
}
