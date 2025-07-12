package models

import "time"

type TransactionID int64
type UserID int64

// Transaction — основная сущность транзакции
type Transaction struct {
    ID           TransactionID                      // Уникальный идентификатор транзакции
    CreatorID    UserID                      // ID пользователя, создавшего транзакцию
    Amount       int64                      // Сумма транзакции (например, в копейках)
    Description  string                     // Описание (например, "Ужин в кафе")
    CreatedAt    time.Time                  // Время создания
    Participants []TransactionParticipant   // Список участников
}

// TransactionParticipant — участник транзакции
type TransactionParticipant struct {
    UserID UserID // ID пользователя-участника
    Amount int64 // Сколько должен этот участник (или переплатил)
}