package api

import (
	"net/http"

	"balance/internal/handlers"
	"balance/internal/repository"
	"balance/internal/service"

	"github.com/jackc/pgx/v5"
)

func NewRouter(db *pgx.Conn) http.Handler {
	// Создаем репозитории
	userRepo := repository.NewUserRepository(db)

	// Создаем сервисы
	userService := service.NewUserService(userRepo)

	// Создаем обработчики
	userHandler := handlers.NewUserHandler(userService)

	// Создаем мультиплексор
	mux := http.NewServeMux()

	// API v1 маршруты
	apiV1 := http.NewServeMux()

	// Пользователи
	apiV1.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.GetAllUsers(w, r)
		case http.MethodPost:
			userHandler.CreateUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	apiV1.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.GetUser(w, r)
		case http.MethodPut:
			userHandler.UpdateUser(w, r)
		case http.MethodDelete:
			userHandler.DeleteUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Добавляем CORS middleware
	mux.Handle("/api/v1/", corsMiddleware(apiV1))

	// Health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","message":"Server is running"}`))
	})

	return mux
}

// corsMiddleware добавляет CORS заголовки
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Разрешаем запросы с любого origin (для разработки)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "86400")

		// Обрабатываем preflight запросы
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
