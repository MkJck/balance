package services

import (
    "errors"
    "github.com/MkJck/balance/internal/models"
    "github.com/MkJck/balance/internal/repository"
)

// TransactionService — сервис для работы с транзакциями
type TransactionService struct {
    repo repository.TransactionRepository
}

// Конструктор сервиса
func NewTransactionService(repo repository.TransactionRepository) *TransactionService {
    return &TransactionService{repo: repo}
}

// Создание транзакции
func (s *TransactionService) CreateTransaction(
    creatorID models.UserID,
    amount int64,
    participantIDs []models.UserID,
    description string,
) (*models.Transaction, error) {
    if amount <= 0 {
        return nil, errors.New("amount must be positive")
    }
    if len(participantIDs) == 0 {
        return nil, errors.New("participants required")
    }
    share := amount / int64(len(participantIDs))
    participants := make([]models.TransactionParticipant, len(participantIDs))
    for i, uid := range participantIDs {
        participants[i] = models.TransactionParticipant{
            UserID: uid,
            Amount: share,
        }
    }
    tx := &models.Transaction{
        CreatorID:    creatorID,
        Amount:       amount,
        Description:  description,
        Participants: participants,
    }
    id, err := s.repo.Create(tx)
    if err != nil {
        return nil, err
    }
    tx.ID = id
    return tx, nil
}

// Получение транзакции по ID
func (s *TransactionService) GetTransaction(id models.TransactionID) (*models.Transaction, error) {
    return s.repo.GetByID(id)
}

// Получение всех транзакций пользователя
func (s *TransactionService) ListTransactionsByUser(userID models.UserID) ([]*models.Transaction, error) {
    return s.repo.ListByUser(userID)
}