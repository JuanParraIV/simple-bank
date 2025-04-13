DB_URL=postgresql://root:TEst.0429.30@localhost:5432/simple_bank?sslmode=disable
postgres:
	docker run --name simpledb -e POSTGRES_USER=root -e POSTGRES_PASSWORD=TEst.0429.30 -p 5432:5432 -d postgres:14-alpine
	@echo "Postgres container started on port 5432"
	@echo "Postgres container started with name simpledb"

postgresdown:
	docker stop simpledb
	docker rm simpledb
	@echo "Postgres container stopped and removed"

createdb:
	docker exec -it simpledb createdb --username=root --owner=root simple_bank
	@echo "Database simple_bank created"

dropdb:
	docker exec -it simpledb dropdb simple_bank
	@echo "Database simple_bank dropped"

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up
	@echo "Migration up completed"

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down
	@echo "Migration down completed"

sqlc:
	sqlc generate
	@echo "SQLC generated at db/sqlc"

test:
	go test -v -cover ./...
	@echo "Tests completed"

server:
	go run main.go
	@echo "Server started on port 8080"

mock:
	mockgen -package mockdb  -destination db/mock/store.go github.com/juanparraiv/simple-bank/db/sqlc Store
	@echo "Mock generated at db/mock/store.go"

.PHONY: createdb dropdb postgres postgresdown migrateup migratedown sqlc test server mock

