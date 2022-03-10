migrateup:
	migrate -path db/migration -database "postgresql://postgres:password@127.0.0.1/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:password@127.0.0.1/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: migrateup migratedown sqlc
