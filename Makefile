migrateup:
	migrate -path db/migration -database "postgresql://postgres:password@127.0.0.1/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://postgres:password@127.0.0.1/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://postgres:password@127.0.0.1/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://postgres:password@127.0.0.1/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover -vet=all -parallel=4 ./...

server:
	go run main.go

.PHONY: migrateup migrateup1 migratedown migratedown1 sqlc server
