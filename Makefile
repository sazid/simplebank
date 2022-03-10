migrateup:
	migrate -path db/migration -database "postgresql://postgres:password@127.0.0.1/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:password@127.0.0.1/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover -vet=all -parallel=4 ./...

.PHONY: migrateup migratedown sqlc
