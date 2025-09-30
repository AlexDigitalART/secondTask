# Makefile для управления миграциями и приложением

# Переменные (настройки)
DB_DSN := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Создать новую миграцию
migrate-new:
	migrate create -ext sql -dir ./migrations $(NAME)

# Применить миграции (создать таблицы)
migrate-up:
	$(MIGRATE) up

# Откатить миграции (удалить таблицы)
migrate-down:
	$(MIGRATE) down

# Запустить приложение
run:
	go run main.go