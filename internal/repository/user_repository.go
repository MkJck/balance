package repository

import (
	"context"
	"fmt"

	"balance/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *models.CreateUserRequest) (*models.User, error) {
	query := `
		INSERT INTO users (name, email) 
		VALUES ($1, $2) 
		RETURNING id, name, email, created_at, updated_at
	`
	
	var result models.User
	err := r.db.QueryRow(ctx, query, user.Name, user.Email).Scan(
		&result.ID, &result.Name, &result.Email, &result.CreatedAt, &result.UpdatedAt,
	)
	
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return nil, fmt.Errorf("email already exists")
		}
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	
	return &result, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
	query := `
		SELECT id, name, email, created_at, updated_at 
		FROM users 
		WHERE id = $1
	`
	
	var user models.User
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt,
	)
	
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	
	return &user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT id, name, email, created_at, updated_at 
		FROM users 
		WHERE email = $1
	`
	
	var user models.User
	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt,
	)
	
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	
	return &user, nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*models.User, error) {
	query := `
		SELECT id, name, email, created_at, updated_at 
		FROM users 
		ORDER BY created_at DESC
	`
	
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	defer rows.Close()
	
	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, &user)
	}
	
	return users, nil
}

func (r *UserRepository) Update(ctx context.Context, id int, user *models.UpdateUserRequest) (*models.User, error) {
	query := `
		UPDATE users 
		SET name = $1, email = $2, updated_at = CURRENT_TIMESTAMP 
		WHERE id = $3 
		RETURNING id, name, email, created_at, updated_at
	`
	
	var result models.User
	err := r.db.QueryRow(ctx, query, user.Name, user.Email, id).Scan(
		&result.ID, &result.Name, &result.Email, &result.CreatedAt, &result.UpdatedAt,
	)
	
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return nil, fmt.Errorf("email already exists")
		}
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	
	return &result, nil
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	
	if result.RowsAffected() == 0 {
		return fmt.Errorf("user not found")
	}
	
	return nil
} 