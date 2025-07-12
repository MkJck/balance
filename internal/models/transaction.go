package models

import "time"

type Transaction struct {
    ID          int64
    CreatorID   int64
    Amount      int64 // в рублях
    Description string
    CreatedAt   time.Time
    Participants []TransactionParticipant
}

type TransactionParticipant struct {
    UserID int64
    Amount int64 // сколько должен этот участник
}