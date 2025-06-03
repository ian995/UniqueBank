postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17.5-alpine3.21
createdb:
	docker exec -it postgres17 createdb --username=root --owner=root uniquebank

dropdb:
	docker exec -it postgres17 dropdb --username=root --if-exists uniquebank

migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/uniquebank?sslmode=disable" -verbose up
	
migrateup1:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/uniquebank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/uniquebank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/uniquebank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate -f config/sqlc.yaml

test:
	go test ./... -v -coverpkg=./...

server:
	go run cmd/server/main.go

mock:
	mockgen -source=internal/repo/store.go -package mock_test -aux_files=github.com/ian995/UniqueBank/internal/repo=internal/repo/querier.go  -destination=tests/mock/mock_db.go

.phony: postgres createdb dropdb migrateup migratedown sqlc test server migratedown1 migrateup1