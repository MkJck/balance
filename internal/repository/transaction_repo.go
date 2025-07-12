package repository

import "github.com/MkJck/balance/internal/models"

// TransactionRepository — интерфейс для работы с транзакциями
type TransactionRepository interface {
    Create(tx *models.Transaction) (models.TransactionID, error)           // Создать транзакцию, вернуть её ID
    GetByID(id models.TransactionID) (*models.Transaction, error)          // Получить транзакцию по ID
    ListByUser(userID models.UserID) ([]*models.Transaction, error)		// Получить все транзакции пользователя
    // Можно добавить другие методы по необходимости
}