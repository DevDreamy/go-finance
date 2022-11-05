createDb: 
	createdb --username-postgres --owner=postgres go_finance

migrateUp: 
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable" -verbose up

migrateDrop: 
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable" -verbose drop

.PHONY: createDb migrateUp migrateDrop