package service

import (
	"context"
	"fmt"

	"balance/internal/models"
	"balance/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.User, error) {
	// Валидация данных
	if err := s.validateCreateUserRequest(req); err != nil {
		return nil, err
	}

	// Проверяем, не существует ли уже пользователь с таким email
	existingUser, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("user with this email already exists")
	}

	// Создаем пользователя
	user, err := s.userRepo.Create(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

func (s *UserService) GetUser(ctx context.Context, id int) (*models.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	users, err := s.userRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	return users, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int, req *models.UpdateUserRequest) (*models.User, error) {
	// Проверяем, существует ли пользователь
	existingUser, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	// Если email изменился, проверяем уникальность
	if req.Email != "" && req.Email != existingUser.Email {
		userWithEmail, err := s.userRepo.GetByEmail(ctx, req.Email)
		if err == nil && userWithEmail != nil {
			return nil, fmt.Errorf("user with this email already exists")
		}
	}

	// Обновляем пользователя
	user, err := s.userRepo.Update(ctx, id, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	// Проверяем, существует ли пользователь
	_, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	// Удаляем пользователя
	if err := s.userRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

func (s *UserService) validateCreateUserRequest(req *models.CreateUserRequest) error {
	if req.Name == "" {
		return fmt.Errorf("name is required")
	}
	if len(req.Name) < 2 {
		return fmt.Errorf("name must be at least 2 characters long")
	}
	if len(req.Name) > 100 {
		return fmt.Errorf("name must be no more than 100 characters long")
	}

	if req.Email == "" {
		return fmt.Errorf("email is required")
	}

	// Простая валидация email
	if !s.isValidEmail(req.Email) {
		return fmt.Errorf("invalid email format")
	}

	return nil
}

func (s *UserService) isValidEmail(email string) bool {
	// Простая проверка email - в реальном проекте лучше использовать библиотеку
	if len(email) < 5 {
		return false
	}

	hasAt := false
	hasDot := false

	for i, char := range email {
		if char == '@' {
			if hasAt || i == 0 || i == len(email)-1 {
				return false
			}
			hasAt = true
		}
		if char == '.' && hasAt {
			hasDot = true
		}
	}

	return hasAt && hasDot
}
