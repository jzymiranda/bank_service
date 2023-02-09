
postgres:
	docker run --name postgres12 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine


createdb:
	docker exec -it postgres12 createdb --username=root --owner=root bank_service

dropdb:
	docker exec -it postgres12 dropdb bank_service


migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/bank_service?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/bank_service?sslmode=disable" -verbose up 1


migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/bank_service?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/bank_service?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -build_flags=--mod=mod -package mockdb -destination db/mock/store.go github.com/jzymiranda/bank_service/db/output Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server mock