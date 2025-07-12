package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/MkJck/balance/internal/models"
    "github.com/MkJck/balance/internal/services"
)

// TransactionHandler — структура для работы с транзакциями через HTTP
type TransactionHandler struct {
    service *services.TransactionService
}

// Конструктор
func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
    return &TransactionHandler{service: service}
}

// Структура для запроса создания транзакции
type createTransactionRequest struct {
    Amount	        int64           `json:"amount"`
    ParticipantIDs	[]models.UserID `json:"participants"`
    Description   	string          `json:"description"`
}

// POST /transactions
func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
    var req createTransactionRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // В реальном проекте creatorID берём из аутентификации, здесь — заглушка
    creatorID := models.UserID(1)
    tx, err := h.service.CreateTransaction(creatorID, req.Amount, req.ParticipantIDs, req.Description)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tx)
}