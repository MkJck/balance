package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"balance/api"
	"balance/internal/config"
	"balance/internal/database"

	"github.com/joho/godotenv"
)

func main() {
	// Загружаем .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Инициализируем конфигурацию
	cfg := config.New()

	// Подключаемся к базе данных
	db, err := database.Connect(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close(context.Background())

	// Применяем миграции
	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Создаем роутер
	router := api.NewRouter(db)

	// Создаем HTTP сервер
	server := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Канал для graceful shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Запускаем сервер в горутине
	go func() {
		log.Printf("Server is running on: http://localhost:%s", cfg.Server.Port)
		log.Println("Available endpoints:")
		log.Println("GET  /api/v1/users")
		log.Println("POST /api/v1/users")
		log.Println("GET  /api/v1/debts")
		log.Println("POST /api/v1/debts")
		log.Println("GET  /api/v1/groups")
		log.Println("POST /api/v1/groups")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Ждем сигнал для graceful shutdown
	<-done
	log.Println("Server is shutting down...")

	// Даем серверу время на завершение текущих запросов
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}

	log.Println("Server stopped")
}
