package main

import (
    "github.com/gin-gonic/gin"
    "github.com/MkJck/balance/internal/handlers"
    "github.com/MkJck/balance/internal/services"
    "github.com/MkJck/balance/internal/repository"
    "github.com/gin-contrib/cors"
    "time"
    // "github.com/joho/godotenv"
    // "os"
    // Здесь можно добавить импорт реальной реализации репозитория, например, для PostgreSQL
)

func main() {

    // err := godotenv.Load()
    // if err != nil {
    //     log.Fatal("Error loading .env file")
    // }
    // dbHost := os.Getenv("DB_HOST")

    

    // 1. Создаём репозиторий (пока можно использовать заглушку или in-memory)
    var txRepo repository.TransactionRepository
    // TODO: Реализуй и подключи реальный репозиторий (например, PostgreSQL)
    // txRepo = postgres.NewTransactionRepo(...)

    // 2. Создаём сервис
    txService := services.NewTransactionService(txRepo)

    // 3. Создаём обработчик
    txHandler := handlers.NewTransactionHandler(txService)

    // 4. Создаём роутер Gin
    r := gin.Default()

    // Настройка CORS
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, // URL твоего фронтенда
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    // 5. Регистрируем маршруты
    r.POST("/transactions", txHandler.CreateTransaction)
    // Можно добавить другие маршруты: r.GET("/transactions", ...)

    // 6. Запускаем сервер
    r.Run(":8080") // По умолчанию порт 8080
}