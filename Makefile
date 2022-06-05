postgres:
	docker run --name test-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mypassword -d postgres:12-alpine
createdb:
	docker exec -it test-postgres createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it test-postgres dropdb simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:mypassword@localhost:5432/simple_bank?sslmode=disable" -verbose up 
migratedown:
	migrate -path db/migration -database "postgresql://root:mypassword@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
.PHONY: postgres createdb dropdb migrateup migratedown sqlc	test