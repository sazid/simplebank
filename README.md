# Simple Bank

A toy project to learn about production coding.

## Docker

1. Build docker image: `docker build -t simplebank:latest .`

2. Create a network for running both the PostgreSQL container and the app binary
   they can discover each other: `docker network create bank-network`

3. Run docker container: `docker run --name simplebank --network bank-network -p 8300:8300 -e GIN_MODE="release" -e DB_SOURCE="postgresql://postgres:password@db:5432/simple_bank?sslmode=disable" simplebank:latest`

## How to generate code

- Generate SQL CRUD with sqlc:

   ```bash
   make sqlc
   ```

- Create a new DB migration:

   ```bash
   migrate create -ext sql -dir db/migration -seq <migration_name>
   ```

## How to run

- Run server:

   ```bash
   make server
   ```
