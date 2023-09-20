postgres:
	docker run --name postgres12 -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -p 5439:5432 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate --path db/migrations --database "postgresql://root:secret@localhost:5439/simple_bank?sslmode=disable" --verbose up

migratedown:
	migrate --path db/migrations --database "postgresql://root:secret@localhost:5439/simple_bank?sslmode=disable" --verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./... 

.PHONY: postgres createdb migrateup migratedown sqlc test