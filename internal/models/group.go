package models

import (
	"time"
)

type Group struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedBy   int       `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Дополнительные поля для API
	CreatedByUser *User   `json:"created_by_user,omitempty"`
	Members       []*User `json:"members,omitempty"`
	Debts         []*Debt `json:"debts,omitempty"`
}

type CreateGroupRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=100"`
	Description string `json:"description" validate:"max=500"`
	CreatedBy   int    `json:"created_by" validate:"required"`
}

type UpdateGroupRequest struct {
	Name        *string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Description *string `json:"description,omitempty" validate:"omitempty,max=500"`
}

type AddMemberRequest struct {
	UserID int `json:"user_id" validate:"required"`
}

type GroupSummary struct {
	GroupID      int     `json:"group_id"`
	GroupName    string  `json:"group_name"`
	TotalDebts   float64 `json:"total_debts"`
	ActiveDebts  int     `json:"active_debts"`
	SettledDebts int     `json:"settled_debts"`
	MemberCount  int     `json:"member_count"`
}
