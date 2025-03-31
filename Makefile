DB_URL=postgresql://root:TEst.0429.30@localhost:5433/simple_bank?sslmode=disable
postgres:
	docker run --name simpledb -e POSTGRES_USER=root -e POSTGRES_PASSWORD=TEst.0429.30 -p 5433:5432 -d postgres:14-alpine

postgresdown:
	docker stop simpledb
	docker rm simpledb
createdb:
	docker exec -it simpledb createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it simpledb dropdb simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up
migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...
.PHONY: createdb dropdb postgres postgresdown migrateup migratedown sqlc test

