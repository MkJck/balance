package models

import (
	"time"
)

type Debt struct {
	ID          int        `json:"id"`
	GroupID     int        `json:"group_id"`
	FromUserID  int        `json:"from_user_id"`
	ToUserID    int        `json:"to_user_id"`
	Amount      float64    `json:"amount"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	SettledAt   *time.Time `json:"settled_at,omitempty"`

	// Дополнительные поля для API
	FromUser *User  `json:"from_user,omitempty"`
	ToUser   *User  `json:"to_user,omitempty"`
	Group    *Group `json:"group,omitempty"`
}

type CreateDebtRequest struct {
	GroupID     int     `json:"group_id" validate:"required"`
	FromUserID  int     `json:"from_user_id" validate:"required"`
	ToUserID    int     `json:"to_user_id" validate:"required"`
	Amount      float64 `json:"amount" validate:"required,gt=0"`
	Description string  `json:"description" validate:"max=500"`
}

type UpdateDebtRequest struct {
	Amount      *float64 `json:"amount,omitempty" validate:"omitempty,gt=0"`
	Description *string  `json:"description,omitempty" validate:"omitempty,max=500"`
	Status      *string  `json:"status,omitempty" validate:"omitempty,oneof=active settled cancelled"`
}

type DebtSummary struct {
	UserID      int     `json:"user_id"`
	UserName    string  `json:"user_name"`
	TotalOwed   float64 `json:"total_owed"`    // Сколько должен пользователь
	TotalOwedTo float64 `json:"total_owed_to"` // Сколько должны пользователю
	NetBalance  float64 `json:"net_balance"`   // Чистый баланс (положительный = должны ему, отрицательный = он должен)
}
