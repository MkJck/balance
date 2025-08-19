# Balance - Приложение для отслеживания долгов между друзьями

## Описание

Backend API для приложения по отслеживанию долгов между друзьями. Позволяет создавать группы друзей, добавлять долги и отслеживать балансы.

## Архитектура

Проект организован по принципам Clean Architecture:

```
balance/
├── cmd/server/          # Точка входа приложения
├── internal/            # Внутренняя логика приложения
│   ├── config/         # Конфигурация
│   ├── database/       # Работа с базой данных
│   ├── handlers/       # HTTP обработчики
│   ├── models/         # Модели данных
│   ├── repository/     # Слой доступа к данным
│   └── service/        # Бизнес-логика
├── pkg/                # Переиспользуемые пакеты
│   └── utils/          # Утилиты
└── api/                # API роутинг
```

## Установка и запуск

### Требования

- Go 1.21+
- PostgreSQL 12+

### Настройка

1. Клонируйте репозиторий:
```bash
git clone <repository-url>
cd balance
```

2. Установите зависимости:
```bash
go mod tidy
```

3. Создайте файл `.env` на основе `.env.example`:
```bash
# Server Configuration
SERVER_PORT=8080

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=your_password
DB_NAME=balance_db
DB_SSLMODE=disable
```

4. Создайте базу данных PostgreSQL:
```sql
CREATE DATABASE balance_db;
```

5. Запустите приложение:
```bash
go run cmd/server/main.go
```

## API Endpoints

### Пользователи

- `GET /api/v1/users` - Получить всех пользователей
- `POST /api/v1/users` - Создать пользователя
- `GET /api/v1/users?id=1` - Получить пользователя по ID
- `PUT /api/v1/users?id=1` - Обновить пользователя
- `DELETE /api/v1/users?id=1` - Удалить пользователя

### Примеры запросов

#### Создание пользователя
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Иван Иванов",
    "email": "ivan@example.com"
  }'
```

#### Получение всех пользователей
```bash
curl http://localhost:8080/api/v1/users
```

## Структура базы данных

### Таблицы

- `users` - Пользователи
- `groups` - Группы друзей
- `group_members` - Участники групп
- `debts` - Долги

### Схема

```sql
-- Пользователи
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Группы
CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    created_by INTEGER REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Участники групп
CREATE TABLE group_members (
    id SERIAL PRIMARY KEY,
    group_id INTEGER REFERENCES groups(id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(group_id, user_id)
);

-- Долги
CREATE TABLE debts (
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
);
```

## Разработка

### Добавление новых функций

1. Создайте модель в `internal/models/`
2. Создайте репозиторий в `internal/repository/`
3. Создайте сервис в `internal/service/`
4. Создайте обработчик в `internal/handlers/`
5. Добавьте маршруты в `api/routes.go`

### Тестирование

```bash
go test ./...
```

## Планы развития

- [ ] Добавить аутентификацию и авторизацию
- [ ] Реализовать API для групп
- [ ] Реализовать API для долгов
- [ ] Добавить уведомления
- [ ] Создать фронтенд приложение
- [ ] Добавить тесты
- [ ] Добавить документацию API (Swagger) 