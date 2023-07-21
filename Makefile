run: 
	go run cmd/main.go

dbup:
	migrate -path pkg/db/orm/migration -database "postgres://root:password@localhost:5432/dda?sslmode=disable" -verbose up

dbdown:
	migrate -path pkg/db/orm/migration -database "postgres://root:password@localhost:5432/dda?sslmode=disable" -verbose down

generate:
	sqlc generate

.PHONY: generate