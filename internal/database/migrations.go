package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func RunMigrations(db *pgx.Conn) error {
	ctx := context.Background()

	migrations := []string{
		// Таблица пользователей
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,

		// Таблица групп друзей
		`CREATE TABLE IF NOT EXISTS groups (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT,
			created_by INTEGER REFERENCES users(id),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,

		// Таблица участников групп
		`CREATE TABLE IF NOT EXISTS group_members (
			id SERIAL PRIMARY KEY,
			group_id INTEGER REFERENCES groups(id) ON DELETE CASCADE,
			user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
			joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			UNIQUE(group_id, user_id)
		)`,

		// Таблица долгов
		`CREATE TABLE IF NOT EXISTS debts (
			id SERIAL PRIMARY KEY,
			group_id INTEGER REFERENCES groups(id) ON DELETE CASCADE,
			from_user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
			to_user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
			amount DECIMAL(10,2) NOT NULL CHECK (amount > 0),
			description TEXT,
			status TEXT DEFAULT 'active' CHECK (status IN ('active', 'settled', 'cancelled')),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			settled_at TIMESTAMP,
			CHECK (from_user_id != to_user_id)
		)`,

		// Индексы для оптимизации
		`CREATE INDEX IF NOT EXISTS idx_debts_group_id ON debts(group_id)`,
		`CREATE INDEX IF NOT EXISTS idx_debts_from_user_id ON debts(from_user_id)`,
		`CREATE INDEX IF NOT EXISTS idx_debts_to_user_id ON debts(to_user_id)`,
		`CREATE INDEX IF NOT EXISTS idx_debts_status ON debts(status)`,
		`CREATE INDEX IF NOT EXISTS idx_group_members_group_id ON group_members(group_id)`,
		`CREATE INDEX IF NOT EXISTS idx_group_members_user_id ON group_members(user_id)`,
	}

	for i, migration := range migrations {
		if _, err := db.Exec(ctx, migration); err != nil {
			log.Printf("Migration %d failed: %v", i+1, err)
			return err
		}
		log.Printf("Migration %d completed successfully", i+1)
	}

	log.Println("All migrations completed successfully")
	return nil
}
