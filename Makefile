postgresdb:
	docker run --name postgresdb -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=saya -d postgres:16.2-alpine3.19

createdb:
	docker exec -it postgresdb createdb --username=root --owner=root school_management

dropdb:
	docker exec -it postgresdb dropdb school_management

migrateup:
	migrate -path db/migration -database "postgresql://root:saya@localhost:5432/school_management?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:saya@localhost:5432/school_management?sslmode=disable" -verbose down

migrateforce:
	migrate -path db/migration -database "postgresql://root:saya@localhost:5432/school_management?sslmode=disable" -verbose force 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgresdb createdb dropdb sqlc test