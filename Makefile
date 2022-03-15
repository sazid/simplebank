DB_URL=postgresql://root:password@localhost:5432/simple_bank?sslmode=disable

network:
	docker network create bank-network

postgres:
	docker run --name postgres11 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:11-alpine

createdb:
	docker exec -it postgres11 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres11 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover -vet=all -parallel=4 ./...

server:
	go run main.go

.PHONY: migrateup migrateup1 migratedown migratedown1 sqlc server postgres network
