postgres:
	docker run --name quranpostgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it quranpostgres createdb --username=root --owner=root quran_db

dropdb:
	docker exec -it quranpostgres dropdb quran_db

migratecreate:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/quran_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/quran_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

oapi:
	oapi-codegen -generate types -o ./api/types.gen.go -package api ./doc/openapi.yaml
	oapi-codegen -generate gin-server -o ./api/server.gen.go -package api ./doc/openapi.yaml

server:
	kill $(lsof -i :9090 -t) &
	nohup swagger-ui -p=9090 ./doc/openapi.yaml &> /dev/null &
	go run main.go

scrap:
	go run ./scripts/scrap_data.go

.PHONY: postgres createdb dropdb migratecreate migrateup migratedown sqlc test server