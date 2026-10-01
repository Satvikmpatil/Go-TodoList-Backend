DB_URL=postgresql://postgres:mypassword@localhost:5500/postgres?sslmode=disable

postgres:
	docker run -d --name my-postgres -e POSTGRES_PASSWORD=mypassword -p 5500:5432 postgres:14-alpine

createdb:
	docker exec -it my-postgres createdb --username=postgres --owner=postgres todolist

dropdb:
	docker exec -it my-postgres dropdb --username=postgres todolist

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./db/sqlc/...

.PHONY: postgres createdb dropdb migrateup migratedown migrateforce test sqlc
