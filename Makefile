# Makefile для управления миграциями и приложением

# Переменные (настройки)
DB_DSN := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/api.gen.go

lint:
	golangci-lint run --color=auto

# Создать новую миграцию
migrate-new:
	migrate create -ext sql -dir ./migrations $(NAME)

# Применить миграции (создать таблицы)
migrate-up:
	$(MIGRATE) up

# Откатить миграции (удалить таблицы)
migrate-down:
	$(MIGRATE) down

create-migration-users:
	migrate create -ext sql -digits 6 -dir ./migrations -seq create_users_table

# Запустить приложение
run:
	go run main.go

test:
	go test ./... -v