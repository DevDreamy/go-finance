createdb:
	createdb --username-postgres --owner=postgres go_finance

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable" -verbose up

migratedrop:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable" -verbose drop

test:
	go test -v -cover ./...

.PHONY: createdb migrateup migratedrop test